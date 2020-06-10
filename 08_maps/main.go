package main

import "fmt"

func main() {
	emails := make(map[string]string)

	// Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["Mike"] = "Mike@gmail.com"

	fmt.Println(emails)

	fmt.Println(emails["Bob"])

	// Delete from maps
	delete(emails, "Bob")
	fmt.Println(emails)

	capitals := map[string]string{"India": "New Delhi", "USA": "Washington DC"}
	fmt.Println(capitals)
	fmt.Println(len(capitals))

}
