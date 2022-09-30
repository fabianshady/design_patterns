package main

import "fmt"

func main() {
	hondaFactory, _ := GetVehicleFactory("honda")
	heroFactory, _ := GetVehicleFactory("hero")

	hondaBike := hondaFactory.GetBike()
	hondaScooter := hondaFactory.GetScooter()

	heroBike := heroFactory.GetBike()
	heroScooter := heroFactory.GetScooter()
}