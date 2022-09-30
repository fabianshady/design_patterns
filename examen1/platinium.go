package main

type Platinium struct {
	CreditCard
}

func newPlatinium() ICreditCard {
	return &Platinium{
		CreditCard: CreditCard{
			cardtype:    "Platinium",
			creditlimit: 40000,
			cat:         59,
		},
	}
}
