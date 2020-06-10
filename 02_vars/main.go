package main

import "fmt"

var name string = "Vivek Chandra"

func main() {
	// using var keyword
	name := "Vivek Chandra BS"
	var age int32 = 37
	var isCool = true
	size := 1.3
	email := ""

	fmt.Println(name)
	fmt.Println(age)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)

}
