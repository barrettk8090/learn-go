package main

import "fmt"

func main() {
	// strings
	var nameOne string = "bob"
	var nameTwo = "sam"
	var nameThree string
	fmt.Println("On first run: ", nameOne, nameTwo, nameThree)

	nameOne = "derek"
	nameThree = "alfred"

	fmt.Println("On second run:", nameOne, nameTwo, nameThree)

	// shorthand - can only use inside a function
	nameFour := "sally"
	fmt.Println("Name Four: ", nameFour)
}
