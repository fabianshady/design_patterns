package main

type Pizza struct {
	ingredient string
	price      float64
}

func (g *Pizza) setIngredient(ingredient string) {
	g.ingredient = ingredient
}

func (g *Pizza) getIngredient() string {
	return g.ingredient
}

func (g *Pizza) setPrice(price float64) {
	g.price = price
}

func (g *Pizza) getPrice() float64 {
	return g.price
}
