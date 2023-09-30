package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bill

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	} //declaration and initialization
	return b
}

//Receiver function

func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip) //-25 make the value will be printed 25 characters long and add spaces if the string is less than 25 (to the right) (+25 to the left)

	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total+b.tip)

	return fs
}

// update bill
func (b *bill) updatedTip(tip float64) {
	b.tip = tip //dereffrence automatically
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

//SAVE BILL

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}
