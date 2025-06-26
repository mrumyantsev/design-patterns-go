package main

import (
	"fmt"
)

func main() {
	var beverage1 Beverager = NewEspresso()

	printDescriptionAndCost(beverage1)

	var beverage2 Beverager = NewDarkRoast()
	beverage2 = NewMocha(beverage2)
	beverage2 = NewMocha(beverage2)
	beverage2 = NewWhip(beverage2)

	printDescriptionAndCost(beverage2)

	var beverage3 Beverager = NewHouseBlend()
	beverage3 = NewSoy(beverage3)
	beverage3 = NewMocha(beverage3)
	beverage3 = NewWhip(beverage3)

	printDescriptionAndCost(beverage3)
}

func printDescriptionAndCost(beverage Beverager) {
	fmt.Printf("%s $%.2f\n", beverage.Description(), beverage.Cost())
}
