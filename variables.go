package main

import "fmt"

func variables() {
	//String
	var name1 string = "Itadori"
	var name2 = "Gojo"
	var name3 string
	fmt.Println(name1, name2, name3)

	//we can change values
	name1 = "Gito"
	name3 = "Migumi"
	fmt.Println(name1, name2, name3)

	name4 := "Panda"
	fmt.Println(name4)

	//integers
	var age1 int = 20
	var age2 = 30
	age3 := 40
	fmt.Println(age1, age2, age3)

	//bits and memory
	var num1 int8 = 25 //there is a spesific range of numbers
	var num2 int8 = -128
	var num3 uint8 = 45 //unsigned int
	fmt.Println(num1, num2, num3)

	//Float
	var score1 float32 = -1.5
	var score2 float64 = 123485766675 //we use this at the most time
	score3 := 24.8                    //defined as float64
	fmt.Println(score1, score2, score3)
}
