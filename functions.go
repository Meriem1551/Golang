package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Println("Good morning", n)
}

func sayBye(n string) {
	fmt.Println("Goodbye", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}
