package main

import "fmt"

func main() {
	ids := []int{33, 23, 34, 55, 11}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	capitals := map[string]string{"India": "New Delhi", "USA": "Washington DC"}
	for k, v := range capitals {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range capitals {
		fmt.Println("Name: " + k)
	}

}
