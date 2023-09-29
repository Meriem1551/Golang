package main

import "fmt"

func Arrays_Slices() {
	//Arrays
	//var ages [3]int = [3]int{20, 30, 40}
	var ages = [3]int{20, 30, 40}

	names := [4]string{"Gojo", "Itadori", "Migumi", "Yuta"}
	names[1] = "Gito"
	fmt.Println(ages, len(ages)) //len(arrayName) => length of the array
	fmt.Println(names)

	//Slices {use arrays under the hood}
	var scores = []int{100, 40, 30} // we didn't add number of items
	scores = append(scores, 45)     // return a new slice after adding the new item
	fmt.Println(scores, len(scores))

	//Slice ranges
	range1 := names[1:3] //return from pos 1 to 3 but not including 3 so item 1 and 2
	range2 := names[2:]  // from 2 to the end including the last item
	range3 := names[:3]  // from the first to pos 3 but not including 3
	fmt.Println(range1, range2, range3)
	range1 = append(range1, "Itadori") // add Itadori to the end of range1
	fmt.Println(range1)
}
