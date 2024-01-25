package models

import (
	"fmt"
	"time"
)

const (
	NumberofBarbers         = 2
	NumberofChairs          = 6
	ShopclosingTime         = 20 * time.Second //seconds
	CustomerArrivalInterval = 2                // seconds
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
	fmt.Printf("%d is cutting %d's hair \n", barber, customer)
	time.Sleep(5 * time.Second)
	fmt.Printf("%d is finished cutting %d's hair. \n", barber, customer)
}

func (shop *Shop) Barber(isSleeping bool, barberId int) {
	for {
		if len(shop.WaitingRoom) == 0 {
			fmt.Printf("No customer is in the waiting room and the barber %d goes to sleep\n", barberId)
			isSleeping = true
		}
		customer := <-shop.WaitingRoom
		if shop.IsShopOpen {
			if isSleeping {
				fmt.Printf("%d wakes %d up \n", customer, barberId)
				isSleeping = false
			}
			shop.HairCut(barberId, customer)
		} else {
			fmt.Println("The shop is closed, so the barber has to be sent home")
			return
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
				fmt.Printf("Client %d arrived. %d clients in the waiting room.\n", clientID, len(shop.WaitingRoom))
				clientID++
			} else {
				fmt.Printf("Client %d arrived but the waiting room is full. Client leaves.\n", clientID)
				clientID++
			}
		}
		if !shop.IsShopOpen {
			close(shop.WaitingRoom)
			fmt.Println("the shop is closed, and the client leaves")
			break
		}
	}
	return
}
