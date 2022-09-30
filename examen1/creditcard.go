package main

type CreditCard struct {
	cardtype    string
	creditlimit int
	cat         int
}

func (c *CreditCard) setCardType(cardtype string) {
	c.cardtype = cardtype
}

func (c *CreditCard) getCardType() string {
	return c.cardtype
}

func (c *CreditCard) setCreditLimit(creditlimit int) {
	c.creditlimit = creditlimit
}

func (c *CreditCard) getCreditLimit() int {
	return c.creditlimit
}

func (c *CreditCard) setCAT(cat int) {
	c.cat = cat
}

func (c *CreditCard) getCAT() int {
	return c.cat
}
