package main

type Pepperoni struct {
	Pizza
}

func newPepperoni() IPizza {
	return &Pepperoni{
		Pizza: Pizza{
			ingredient: "Pepperoni",
			price:      150.99,
		},
	}
}
