package models

import (
	"fmt"
	"time"

	"github.com/hariprathap-hp/sleeping_barber_dilemma/utils"
)

const (
	NumberofBarbers         = 1
	NumberofChairs          = 4
	ShopclosingTime         = 20 * time.Second //seconds
	CustomerArrivalInterval = 1                // seconds
)

type Shop struct {
	BarbersEmployed int
	WaitingRoom     chan int
	BarberChan      chan int
	IsShopOpen      bool
}

func NewBarberShop() *Shop {
	return &Shop{
		BarbersEmployed: NumberofBarbers,
		WaitingRoom:     make(chan int, NumberofChairs),
		BarberChan:      make(chan int, NumberofBarbers),
		IsShopOpen:      true,
	}
}

func (shop *Shop) HairCut(barber, customer int) {
	utils.Info(fmt.Sprintf("Barber%d is cutting Client%d's hair \n", barber, customer))
	time.Sleep(5 * time.Second)
	utils.Info(fmt.Sprintf("Barber%d is finished cutting Client%d's hair. \n", barber, customer))
}

func (shop *Shop) Barber(isSleeping bool, barberId int) {
	for {
		if shop.IsShopOpen {
			//If the shop is open and waiting room is empty, the barber goes to sleep
			if len(shop.WaitingRoom) == 0 {
				utils.Info(fmt.Sprintf("No customer is in the waiting room and the barber %d goes to sleep\n", barberId))
				isSleeping = true
			}
			//New customer from waiting room arrives
			customer := <-shop.WaitingRoom
			//If the barber is asleep, the new customer wakes him up, and sets "isSleeping" to false
			if isSleeping {
				utils.Info(fmt.Sprintf("Client%d wakes Barber%d up \n", customer, barberId))
				isSleeping = false
			}
			shop.HairCut(barberId, customer)
		} else {
			if len(shop.WaitingRoom) == 0 {
				utils.Info("The shop is closed, and no one in the waiting room, so the barber has to be sent home")
				return
			} else {
				for customer := range shop.WaitingRoom {
					if isSleeping {
						utils.Info(fmt.Sprintf("Client%d wakes Barber%d up \n", customer, barberId))
						isSleeping = false
					}
					shop.HairCut(barberId, customer)
				}
				utils.Info("The shop is closed, and no one in the waiting room, so the barber has to be sent home 2nd")
				return
			}
		}
	}
}

func (shop *Shop) ClientEntry() {
	fmt.Println("Clients")
	clientID := 1
	for {
		select {
		case <-time.After(CustomerArrivalInterval * time.Second):
			if len(shop.WaitingRoom) < cap(shop.WaitingRoom) {
				shop.WaitingRoom <- clientID
				fmt.Println(len(shop.WaitingRoom), cap(shop.WaitingRoom))
				utils.Info(fmt.Sprintf("Client%d arrived. %d clients in the waiting room.\n", clientID, len(shop.WaitingRoom)))
				clientID++
			} else {
				utils.Info(fmt.Sprintf("Client%d arrived but the waiting room is full. Client leaves.\n", clientID))
				clientID++
			}
		}
		if !shop.IsShopOpen {
			utils.Info("the shop is closed, and the client leaves")
			break
		}
	}
	return
}
