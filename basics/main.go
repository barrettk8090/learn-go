package main

func main() {
	// // strings
	// var nameOne string = "bob"
	// var nameTwo = "sam"
	// var nameThree string
	// fmt.Println("On first run: ", nameOne, nameTwo, nameThree)

	// nameOne = "derek"
	// nameThree = "alfred"

	// fmt.Println("On second run:", nameOne, nameTwo, nameThree)

	// // shorthand - can only use inside a function
	// nameFour := "sally"
	// fmt.Println("Name Four: ", nameFour)

	// // ints
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	// // bits & memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// // uint - positive numbers only, uint8 is 0-255
	// var numThree uint8 = 255

	// fmt.Println(numOne, numTwo, numThree)

	// // floats

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 438573498567.437483
	// scoreThree := 1.5

	// fmt.Println(scoreOne, scoreTwo, scoreThree)

	// // Print
	// fmt.Print("hello, ")
	// fmt.Print("world! \n")
	// fmt.Print("New line \n")

	// // Printf (formatted strings) %_ = format specifier
	// // More info here: https://pkg.go.dev/fmt#hdr-Printing
	// name := "Barrett"
	// age := 30
	// fmt.Printf("my age is %v and my name is %v! \n", age, name)
	// fmt.Printf("age is of type %T \n", age)
	// fmt.Printf("name is of type %T \n", name)
	// fmt.Printf("you scored %f points!\n", 255.55)
	// fmt.Printf("you scored %0.1f points!\n", 255.55)

	// // Sprintf (save formatted strings)
	// var str = fmt.Sprintf("my age is %v and my name is %v! \n", age, name)
	// fmt.Println("The saved string is: ", str)

	// // arrays
	// // var ages [3]int = [3]int{20, 25, 30}
	// ages := [3]int{20, 25, 30}
	// fmt.Println(ages, len(ages))

	// names := [4]string{"yoshi", "mario", "peach", "browser"}
	// names[1] = "luigi"
	// fmt.Println(names, len(names))

	// // slices (use arrays under the hood)

	// var scores = []int{100, 50, 60}
	// scores[2] = 25
	// scores = append(scores, 85)

	// fmt.Println(scores, len(scores))

	// // slice ranges
	// // from position one thru three, but not inclusive of 3
	// rangeOne := names[1:3]

	// // from position two onwards, inclusive of two
	// rangeTwo := names[2:]

	// //  from position 0, up to position 3, but not inclusive of 3
	// rangeThree := names[:3]

	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// rangeOne = append(rangeOne, "koopa")

	// fmt.Println(rangeOne)
}
