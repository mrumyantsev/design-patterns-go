package main

type Soy struct {
	Beverage
}

func NewSoy(beverage Beverager) *Soy {
	return &Soy{
		Beverage: Beverage{
			description: beverage.Description() + ", Soy",
			cost:        beverage.Cost() + 0.15,
		},
	}
}
