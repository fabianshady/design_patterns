package main

import "fmt"

func getCard(creditcardType string) (ICreditCard, error) {
	if creditcardType == "Debit" {
		return newDebit(), nil
	}
	if creditcardType == "Platinium" {
		return newPlatinium(), nil
	}
	if creditcardType == "Titanium" {
		return newTitanium(), nil
	}
	return nil, fmt.Errorf("Wrong card type passed")
}
