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

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	// uint - positive numbers only, uint8 is 0-255
	var numThree uint8 = 255

	fmt.Println(numOne, numTwo, numThree)

	// floats

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 438573498567.437483
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
