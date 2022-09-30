package main

type Titanium struct {
	CreditCard
}

func newTitanium() ICreditCard {
	return &Titanium{
		CreditCard: CreditCard{
			cardtype:    "Titanium",
			creditlimit: 200000,
			cat:         60,
		},
	}
}
