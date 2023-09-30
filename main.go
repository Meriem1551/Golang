package main //every main file must have it
import "fmt"

//import "fmt" //package from go standard library

var score = 86.8

func main() { //we just have one main func
	//variables()
	//fmtMeth()
	//Arrays_Slices()
	//StandardLib()
	//Loops()
	//Bool_and_conditions()

	/*
		sayGreeting("Meriem")
		sayBye("Meriem")
		cycleNames([]string{"Gojo", "Itadori", "Yuta", "Migumi"}, sayGreeting)
		a := circleArea(1.45)
		fmt.Printf("%0.2f", a)
	*/

	/*fn, sn := getInitials("Gojo Gito")
	fmt.Println(fn, sn)*/

	// Package scope
	/*
		sayHello("Meriem")
		for _, v := range points {
			fmt.Println(v)
		}

		showScore()
	*/
	//Map()
	//groupA()
	//groupB()
	//pointers()

	//myBill := newBill("Gojo's bill")

	//myBill.updatedTip(10)
	//myBill.addItem("Sushi", 7.99)
	//myBill.addItem("Ramen", 8.99)
	//myBill.addItem("Curry", 5.50)
	//myBill.addItem("Dorayaki", 1.00)
	//fmt.Println(myBill.format())

	//myBill := createBill()
	//fmt.Println(myBill)
	//promptOptions(myBill)

	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}

}
