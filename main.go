package main

import (
	"sync"
	"time"

	"github.com/hariprathap-hp/sleeping_barber_dilemma/models"
)

func main() {
	shop := models.NewBarberShop()
	var wg sync.WaitGroup

	//Employ the barbers until the shop is closed
	for i := 1; i <= shop.BarbersEmployed; i++ {
		wg.Add(1)
		go func(barberID int) {
			defer wg.Done()
			isSleeping := false
			shop.Barber(isSleeping, barberID)
		}(i)
	}

	//add clients at regular intervals
	wg.Add(1)
	go func() {
		defer wg.Done()
		shop.ClientEntry()
	}()

	//close the shop after certain duration
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-time.After(models.ShopOpenTime)
		shop.IsOpen <- false
		shop.IsShopOpen = false
	}()
	wg.Wait()
}
