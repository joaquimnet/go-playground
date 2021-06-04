package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	menu := map[string]float64{
		"soup":    4.99,
		"spaghet": 7.99,
		"salad":   3.99,
		"yogurt":  5.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	for k, v := range menu {
		fmt.Println(k, "....", v)
	}

	phonebook := map[int]string{
		5550001111: "Kaffe",
		2224449999: "Blu",
		1110003333: "Yoki",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[5550001111])

	phonebook[1110003333] = "Ikaros"
	fmt.Println(phonebook)

}

func otherFiles() {
	// sayHello("Kaffe")

	// for _, v := range points {
	// 	fmt.Println(v)
	// }
}

func multipleReturns() {
	fname, lname := getInitials("blu vidact")
	fmt.Println(fname, lname)

	fname2, lname2 := getInitials("yoki okeya")
	fmt.Println(fname2, lname2)
}

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
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

func funcs1() {
	sayGreeting("Sittle")
	sayGreeting("Kaffe")

	giveItem("Sittle Lister", []string{"Golden Coin", "Shiny Crystal", "Potato", "Key"})
}

func sayGreeting(n string) {
	fmt.Printf("Hello %v!\n", n)
}

func giveItem(name string, items []string) {
	fmt.Printf("%v just got %v items! ", name, len(items))
	for i, item := range items {
		if i == len(items)-1 {
			fmt.Printf("%v.\n", item)
			continue
		}
		fmt.Printf("%v, ", item)
	}
}

func conditionals() {
	age := 22

	if age < 30 {
		fmt.Println("You're young.")
	} else if age < 40 {
		fmt.Println("You're not old.")
	} else {
		fmt.Println("You're getting old. But it's okay!")
	}

	names := []string{"Blu", "Ikaros", "Kaffe"}

	for index, value := range names {
		if index == 1 {
			continue
		}
		fmt.Println(index, value)
	}
}

func looks() {
	x := 0
	for x < 5 {
		fmt.Println("x:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	names := []string{"Blu", "Kaffe", "Ikaros"}

	for i := 0; i < len(names); i++ {
		fmt.Println("name:", names[i])
	}

	for index, val := range names {
		fmt.Printf("Name at %v: %v\n", index, val)
	}
}

func stringsSortAndSearch() {
	greeting := "Hello there fishcake."

	fmt.Println(strings.Contains(greeting, "fishcake"))
	fmt.Println(strings.ReplaceAll(greeting, "l", "w"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "fish"))
	fmt.Println(strings.Split(greeting, " "))

	numbors := []int{34, 56323, 34, 536, 2345, 64, 7, 452, 23, 2, 2, 56, 6, 76, 7}
	sort.Ints(numbors)
	fmt.Println(numbors)

	MINECRAFT_STACK_SIZE := 64

	index := sort.SearchInts(numbors, MINECRAFT_STACK_SIZE)
	fmt.Println(index)

	names := []string{"Blu", "Kaffe", "Ikaros"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Kaffe"))
}

func arrays() {
	var ages [3]int = [3]int{21, 22, 32}

	names := [3]string{"Blu", "Kaffe", "Ikaros"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	scores := []int{69, 420, 666}

	fmt.Println(scores)

	scores = append(scores, 123)

	fmt.Println(scores, len(scores))

	rangeOne := names[1:3]
	rangeTwo := names[1:]
	rangeThree := names[:2]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)
}

func stringsAndLogging() {
	var one string = "Kachow!"
	var two string = "Yeehaw!"
	three := "Fishcake."

	coolNum := 18

	fmt.Println(one, two, three)
	fmt.Println(coolNum)
	fmt.Printf("Some cool vars: %v %v %v \n", one, two, three)

	str := fmt.Sprintf("%v is a cool number", coolNum)
	fmt.Println(">>>", str)
}
