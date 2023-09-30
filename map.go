package main

import "fmt"

func Map() {
	menu := map[string]float64{
		"Ramen":    4.99,
		"Sushi":    7.99,
		"Curry":    6.99,
		"Dorayaki": 3.55,
	}
	fmt.Println(menu)          // all th values and keys
	fmt.Println(menu["sushi"]) // just values

	//Looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type

	phoneBook := map[int]string{
		123456:  "Gojo",
		8285472: "Itadori",
		9837462: "Gito",
	}
	fmt.Println(phoneBook)
	fmt.Println(phoneBook[123456])

	phoneBook[123456] = "Mimi"
	fmt.Println(phoneBook)
}
