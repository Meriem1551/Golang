package main

import "fmt"

func Updatevalue(x *string) {
	*x = "wedge"
}

func pointers() {
	name := "Meriem"
	//fmt.Println("the memory address of name is: ", &name)

	m := &name
	//fmt.Println("the memory address of name is: ", m)
	//fmt.Println("value at memory address: ", *m)
	fmt.Println(name)
	Updatevalue(m)
	fmt.Println(name)
}
