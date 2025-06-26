package main

type Milk struct {
	Beverage
}

func NewMilk(beverage Beverager) *Milk {
	return &Milk{
		Beverage: Beverage{
			description: beverage.Description() + ", Milk",
			cost:        beverage.Cost() + 0.10,
		},
	}
}
