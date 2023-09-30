package main

import "fmt"

func updateName(x string) string {
	// to update the original we need to return x
	x = "wedge"
	return x
}

func groupA() {
	name := "Meriem"
	// updateName(name) //a copy of name
	name = updateName(name) // update the name and change its value
	fmt.Println(name)
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.55
}

func groupB() {
	menu := map[string]float64{
		"Ramen":    4.99,
		"Sushi":    7.99,
		"Curry":    6.99,
		"Dorayaki": 3.55,
	}
	updateMenu(menu) //will change the original value of menu
	fmt.Println(menu)
}
