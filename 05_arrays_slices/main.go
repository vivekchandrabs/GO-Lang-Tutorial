package main

import "fmt"

func main() {
	// Array
	var fruntArr [2]string

	// Assign values
	fruntArr[0] = "Apple"
	fruntArr[1] = "Orange"

	fruitArr2 := [2]string{"Apple", "Orange"}

	fmt.Println(fruntArr)
	fmt.Println(fruntArr[1])
	fmt.Println(fruitArr2)

	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)

	fmt.Println(len(fruitSlice))
	fmt.Print(fruitSlice[1:2])

}
