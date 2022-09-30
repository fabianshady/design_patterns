package main

import "fmt"

type IVehicleFactory interface {
	GetBike() Bike
	GetScooter() Scooter
}

func GetVehicleFactory interface {
	if brand == "honda" {
		return &Honda{}, nil
	}

	if brand == "hero" {
		return &Hero{}, nil
	}

	return nil, fmt.Errorf("Marca erronea")
}