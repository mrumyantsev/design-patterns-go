package main

type Beverage struct {
	description string
	cost        float64
}

func (b *Beverage) Description() string {
	return b.description
}

func (b *Beverage) Cost() float64 {
	return b.cost
}
