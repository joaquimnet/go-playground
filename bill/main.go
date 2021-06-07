package main

import "fmt"

func main() {
	myBill := createBill("Kaffe's Bill")

	myBill.addTip(5.0)
	myBill.addItem("cake", 3.99)
	myBill.addItem("breadsticks", 2.99)
	myBill.addItem("sandwich", 7.99)

	fmt.Println(myBill.format())
}
