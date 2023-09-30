package main

import (
	//"fmt"
	//"math"
	"strings"
)

/*func sayGreeting(n string) {
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
*/
// Return multiple values

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
