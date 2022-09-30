package main

type Pollo struct {
	Pizza
}

func newPollo() IPizza {
	return &Pollo{
		Pizza: Pizza{
			ingredient: "Pollo",
			price:      250.99,
		},
	}
}
