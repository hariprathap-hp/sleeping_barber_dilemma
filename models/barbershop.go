package models

import (
	"fmt"
	"time"
)

const (
	NumberofBarbers         = 3
	NumberofChairs          = 6
	ShopclosingTime         = 21 //seconds
	CustomerArrivalInterval = 3  // seconds
)

type Shop struct {
	barbersEmployed int
	waitingRoom     chan int
	BarberChan      chan int
	isOpen          bool
}

func NewBarberShop() *Shop {
	return &Shop{
		barbersEmployed: NumberofBarbers,
		waitingRoom:     make(chan int, NumberofChairs),
		BarberChan:      make(chan int, NumberofBarbers),
		isOpen:          true,
	}
}

func (shop *Shop) Barber() {
	for i := 0; i < shop.barbersEmployed; i++ {
		shop.BarberChan <- 1
	}
}

func (shop *Shop) Customer() {
	fmt.Println("Customer")
	clientID := 1
	for {
		select {
		case <-time.After(CustomerArrivalInterval * time.Second):
			if len(shop.waitingRoom) < cap(shop.waitingRoom) {
				fmt.Printf("Client %d arrived. %d clients in the waiting room.\n", clientID, len(shop.waitingRoom))
				clientID++
			} else {
				fmt.Printf("Client %d arrived but the waiting room is full. Client leaves.\n", clientID)
				clientID++
			}
		}
	}
}
