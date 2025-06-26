package main

type DarkRoast struct {
	Beverage
}

func NewDarkRoast() *DarkRoast {
	return &DarkRoast{
		Beverage: Beverage{
			description: "Dark Roast Coffee",
			cost:        0.99,
		},
	}
}
