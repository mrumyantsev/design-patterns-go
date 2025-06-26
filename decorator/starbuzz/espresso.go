package main

type Espresso struct {
	Beverage
}

func NewEspresso() *Espresso {
	return &Espresso{
		Beverage: Beverage{
			description: "Espresso",
			cost:        1.99,
		},
	}
}
