package main

import "fmt"

func main() {
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
