package main

import "fmt"

func main() {
	debit, _ := getCard("Debit")
	platinium, _ := getCard("Platinium")
	titanium, _ := getCard("Titanium")

	printDetails(debit)
	printDetails(platinium)
	printDetails(titanium)
}

func printDetails(c ICreditCard) {
	fmt.Printf("Card type: %s", c.getCardType())
	fmt.Println()
	fmt.Printf("Credit limit: %v\n", c.getCreditLimit())
	fmt.Printf("CAT: %v\n", c.getCAT())
	fmt.Println()
}
