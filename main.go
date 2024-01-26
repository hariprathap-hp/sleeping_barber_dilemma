package main

import (
	"fmt"
	"time"

	"github.com/hariprathap-hp/sleeping_barber_dilemma/models"
)

func main() {
	shop := models.NewBarberShop()

	//Employ the barbers until the shop is closed
	for i := 1; i <= shop.BarbersEmployed; i++ {
		isSleeping := false
		go shop.Barber(isSleeping, i)
	}

	//add clients at regular intervals
	go func() {
		shop.ClientEntry()
	}()

	//close the shop after certain duration
	go func() {
		<-time.After(models.ShopclosingTime)
		shop.IsShopOpen = false
		fmt.Println("The shop is closed for the day")
	}()

	time.Sleep(30 * time.Second)
	/*for i := range shop.WaitingRoom {
		fmt.Println("i -- ", i)
	}*/
}
