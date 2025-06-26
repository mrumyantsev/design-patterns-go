package main

type Mocha struct {
	Beverage
}

func NewMocha(beverage Beverager) *Mocha {
	return &Mocha{
		Beverage: Beverage{
			description: beverage.Description() + ", Mocha",
			cost:        beverage.Cost() + 0.20,
		},
	}
}
