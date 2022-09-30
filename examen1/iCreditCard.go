package main

type ICreditCard interface {
	setCardType(cardtype string)
	setCreditLimit(creditlimit int)
	setCAT(cat int)
	getCardType() string
	getCreditLimit() int
	getCAT() int
}
