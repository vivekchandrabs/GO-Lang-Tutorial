package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method ( pointer receover )
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{
		firstName: "Vivek",
		lastName:  "Chandra",
		city:      "Bangalore",
		gender:    "M",
		age:       5}

	person2 := Person{"Bob", "dunn", "New york", "f", 30}

	person1.age++
	person1.getMarried("Gilfoyle")
	person2.getMarried("tom")

	fmt.Println(person1)
	fmt.Println(person1.age)

	person1.hasBirthday()
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
