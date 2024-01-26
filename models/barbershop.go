package models

import (
	"fmt"
	"time"

	"github.com/hariprathap-hp/sleeping_barber_dilemma/utils"
)

const (
	NumberofBarbers         = 2                //Number of barbers in the salon
	WaitingChairs           = 4                //Number of chairs in the waiting room
	ShopOpenTime            = 20 * time.Second //The shop closes in 20 seconds from now
	CustomerArrivalInterval = 1                //The customer arrival interval is set as static. Can be changed to dynamic
	HairCutTime             = 5 * time.Second
)

type Shop struct {
	BarbersEmployed int
	WaitingRoom     chan int
	BarberChan      chan int
	IsShopOpen      bool
	IsOpen          chan bool
}

func NewBarberShop() *Shop {
	return &Shop{
		BarbersEmployed: NumberofBarbers,
		WaitingRoom:     make(chan int, WaitingChairs),
		BarberChan:      make(chan int, NumberofBarbers),
		IsShopOpen:      true,
		IsOpen:          make(chan bool, 1),
	}
}

func (shop *Shop) CloseChannels() {
	utils.Info("Closing all the open channels")
	close(shop.BarberChan)
	close(shop.WaitingRoom)
}

func (shop *Shop) HairCut(barber, customer int) {
	utils.Info(fmt.Sprintf("Barber%d is cutting Client%d's hair \n", barber, customer))
	time.Sleep(HairCutTime)
	utils.Info(fmt.Sprintf("Barber%d is finished cutting Client%d's hair. \n", barber, customer))
}

func (shop *Shop) Barber(isSleeping bool, barberId int) {
	for {
		if shop.IsShopOpen {
			//If the shop is open and waiting room is empty, the barber goes to sleep
			if len(shop.WaitingRoom) == 0 {
				utils.Info(fmt.Sprintf("No customer is in the waiting room and the Barber%d goes to sleep\n", barberId))
				isSleeping = true
			}
			customer := <-shop.WaitingRoom //New customer
			//If the barber is asleep, the new customer wakes him up, and sets "isSleeping" to false
			if isSleeping {
				utils.Info(fmt.Sprintf("Client%d wakes Barber%d up \n", customer, barberId))
				isSleeping = false
			}
			shop.HairCut(barberId, customer)
		} else {
			//If the shop is closed, no new clients should be accepted, but the clients in the waiting room must be attended to
			for len(shop.WaitingRoom) > 0 {
				customer := <-shop.WaitingRoom
				shop.HairCut(barberId, customer)
			}
			//closed the waitingRoom channel once all the customers are served
			utils.Info(fmt.Sprintf("The shop is closed, and no one is in waiting room, so Barber%d can leave the shop", barberId))
			return
		}
	}
}

func (shop *Shop) Customers() {
	clientID := 1
	for {
		select {
		case <-shop.IsOpen:
			utils.Info("The shop is closed, and new clients are not accepted")
			close(shop.IsOpen)
			return
		case <-time.After(CustomerArrivalInterval * time.Second):
			if len(shop.WaitingRoom) < cap(shop.WaitingRoom) {
				shop.WaitingRoom <- clientID
				utils.Info(fmt.Sprintf("Client%d arrived. %d clients in the waiting room.\n", clientID, len(shop.WaitingRoom)))
				fmt.Println(len(shop.WaitingRoom), cap(shop.WaitingRoom))
				clientID++
			} else {
				utils.Info(fmt.Sprintf("Client%d arrives, but the waiting room is full. Client leaves.\n", clientID))
				clientID++
			}
		}
	}
}
