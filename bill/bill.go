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

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format bill
func (b *bill) format() string {
	fs := "Bill breaddown: \n"
	var total float64

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}
	total += b.tip

	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ...$%v\n", "total:", total)

	return fs
}

func (b *bill) addTip(amount float64) {
	b.tip += amount
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) saveBill() {
	data := []byte(b.format())

	err := os.WriteFile("./bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file.")
}
