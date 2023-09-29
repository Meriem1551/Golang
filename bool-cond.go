package main

import "fmt"

func Bool_and_conditions() {
	age := 25
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}
	names := []string{"Gojo", "Itadori", "Migumi", "Yuta"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue //  means don't execute the next instructions and loop again
		}
		if index > 1 {
			fmt.Println("breaking at pos", index)
			break //quit the loop not the program
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
	fmt.Println("end")
}
