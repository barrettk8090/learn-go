package main

import (
	"fmt"
)

func main() {
	// x := 0
	// WHILE LOOP
	// for x < 5 {
	// 	fmt.Println("value of x is: ", x)
	// 	x++
	// }

	// FOR LOOP
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is: ", i)
	// }

	names := []string{"mario", "luigi", "yoshi", "peach"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index, value := range names {
		fmt.Printf("The value at index: %v is %v \n", index, value)
	}

	// value is like a local copy - altering it doesnt alter the value
	// in the original slice.
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
		value = "new string"
	}

	fmt.Println(names)
}
