package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n') // _ mean an error
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) //take input from terminal
	name, _ := getInput("Create a new bill name:", reader)
	b := newBill(name)
	fmt.Println("Create a bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s - save the bill, t - add tip): ", reader)
	fmt.Println(opt)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64) // we parse a string to a float 64 because reader returns strings

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("Item added - ", name, p)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($)", reader)

		t, err := strconv.ParseFloat(tip, 64) // we parse a string to a float 64 because reader returns strings

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updatedTip(t)

		fmt.Println("Tip updated - ", t)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("You saved the file", b.name)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}

}
