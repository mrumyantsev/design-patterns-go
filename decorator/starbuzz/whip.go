package main

type Whip struct {
	Beverage
}

func NewWhip(beverage Beverager) *Whip {
	return &Whip{
		Beverage: Beverage{
			description: beverage.Description() + ", Whip",
			cost:        beverage.Cost() + 0.10,
		},
	}
}
