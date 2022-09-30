package main

type Debit struct {
	CreditCard
}

func newDebit() ICreditCard {
	return &Debit{
		CreditCard: CreditCard{
			cardtype:    "Debit",
			creditlimit: 20000,
			cat:         69,
		},
	}
}
