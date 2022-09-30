package main

import "fmt"

func main() {
	margarita, _ := getPizza("margarita")
	pepperoni, _ := getPizza("pepperoni")
	barbacoa, _ := getPizza("barbacoa")
	pollo, _ := getPizza("pollo")

	printDetails(margarita)
	printDetails(pepperoni)
	printDetails(barbacoa)
	printDetails(pollo)
}

func printDetails(g IPizza) {
	fmt.Printf("Pizza: %s", g.getIngredient())
	fmt.Println()
	fmt.Printf("Price: %v\n", g.getPrice())
	fmt.Println()
}
