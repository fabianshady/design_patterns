package main

type Barbacoa struct {
	Pizza
}

func newBarbacoa() IPizza {
	return &Barbacoa{
		Pizza: Pizza{
			ingredient: "Barbacoa",
			price:      300.99,
		},
	}
}
