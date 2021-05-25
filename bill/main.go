package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
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
