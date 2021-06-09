package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myBill := createBill()
	fmt.Println(myBill)

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput(reader, "Create a new bill name: ")

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)
	promptOptions(b)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput(reader, "Choose option (a - add item, s - save bill, t - add tip): ")
	opt = strings.ToLower(opt)

	switch opt {
	case "a":
		name, _ := getInput(reader, "Item name: ")
		price, _ := getInput(reader, "Item price: ")

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number.")
			promptOptions(b)
			return
		}
		b.addItem(name, p)

		fmt.Println("item added - ", name, p)

		promptOptions(b)
	case "t":
		tip, _ := getInput(reader, "Tip amount ($): ")
		p, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number.")
			promptOptions(b)
			return
		}
		b.addTip(p)

		fmt.Println("tip added - ", p)
		promptOptions(b)
	case "s":
		fmt.Println("You chose S")
	default:
		fmt.Println("That was not a valid option")
		promptOptions(b)
	}
}

func getInput(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	content, err := r.ReadString('\n')
	return strings.TrimSpace(content), err
}
