package main

import (
	"github.com/hariprathap-hp/sleeping_barber_dilemma/models"
)

func main() {
	shop := models.NewBarberShop()

	go func() {
		shop.Customer()
	}()

	/*for i := range shop.BarberChan {
		fmt.Println("i -- ", i)
	}*/
}
