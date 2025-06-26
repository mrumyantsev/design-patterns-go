package main

type Decaf struct {
	Beverage
}

func NewDecaf() *Decaf {
	return &Decaf{
		Beverage: Beverage{
			description: "Decaf Coffee",
			cost:        1.05,
		},
	}
}
