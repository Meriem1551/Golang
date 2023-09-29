package main

import "fmt"

func Loops() {
	x := 0
	for x < 5 {
		//fmt.Println("The value of x is: ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		//fmt.Println("The value of i is: ", i)
	}

	names := [4]string{"Gojo", "Itadori", "Migumi", "Yuta"}
	for i := 0; i < len(names); i++ {
		//fmt.Println(names[i])
	}

	for index, value := range names { // we get the index and the value of slice items
		fmt.Printf("the value at index %v is %v \n", index, value)
		value = "new string"
	}
	fmt.Println(names)
}
