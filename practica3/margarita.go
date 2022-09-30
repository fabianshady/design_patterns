package main

type Margarita struct {
	Pizza
}

func newMargarita() IPizza {
	return &Margarita{
		Pizza: Pizza{
			ingredient: "Margarita",
			price:      200.99,
		},
	}
}
