package main

type IPizza interface {
	setIngredient(ingredient string)
	setPrice(price float64)
	getIngredient() string
	getPrice() float64
}
