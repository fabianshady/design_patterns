package main

import "fmt"

func getPizza(pizzaType string) (IPizza, error) {
	if pizzaType == "margarita" {
		return newMargarita(), nil
	}
	if pizzaType == "pepperoni" {
		return newPepperoni(), nil
	}
	if pizzaType == "barbacoa" {
		return newBarbacoa(), nil
	}
	if pizzaType == "pollo" {
		return newPollo(), nil
	}
	return nil, fmt.Errorf("Wrong pizza type passed")
}
