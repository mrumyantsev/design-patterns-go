package main

type HouseBlend struct {
	Beverage
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{
		Beverage: Beverage{
			description: "House Blend Coffee",
			cost:        0.89,
		},
	}
}
