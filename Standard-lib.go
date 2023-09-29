package main

import (
	"fmt"
	"sort"
	"strings"
)

func StandardLib() {
	//strings package
	greeting := "hello everyone"
	fmt.Println(strings.Contains(greeting, "hello"))         // returns true or false
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) //returns new string with hi instead of hello, it does not change the original one
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll")) //find the pos  where this ll occurs
	fmt.Println(strings.Split(greeting, " "))  //split the string into array when he find a space

	//sort package

	ages := []int{45, 30, 10, 23, 50, 19}
	sort.Ints(ages) //the original slice changed
	fmt.Println(ages)
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"Gojo", "Itadori", "Yuta", "Migumi"}

	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Itadori"))
}
