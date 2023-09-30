package main //every main file must have it
import "fmt"

//import "fmt" //package from go standard library

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

	fn, sn := getInitials("Gojo Gito")
	fmt.Println(fn, sn)
}
