package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// alters the original slice
	sort.Ints(ages)

	fmt.Println(ages)

	// return the position of 30 in ages
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))

}
