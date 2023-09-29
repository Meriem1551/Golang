package main

import "fmt"

func fmtMeth() {
	age := 19
	name := "Meriem"
	//Print
	fmt.Print("Hello, ")
	fmt.Print("World")

	//Println
	fmt.Println("Meriem") //same as print but it add new line automatically
	fmt.Println("Bousaid")
	fmt.Println("my age is: ", age, "and my name is:", name)

	//Printf(Formated strings) %_ = format specifiers
	fmt.Printf("my age is: %v and my name is: %v \n", age, name)
	fmt.Printf("my age is: %q and my name is: %q \n", age, name) //work with strings
	fmt.Printf("age is of type %T", age)                         //type of age
	fmt.Printf("you scored %0.2f points\n", 22.65)

	//Sprintf (save formatted strings)
	str := fmt.Sprintf("my age is %v and my name is %v", age, name) // just save a string
	fmt.Println("The saved string is: ", str)
}
