package main

import (
	"fmt"
	"strings"
)

// n = name
// return two strings
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	// return a slice with the two names separated by space
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	// two or more names? return both initials
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	// single name catch
	return initials[0], "_"
}

func main() {
	// firstname1, secondname1
	fn1, sn1 := getInitials("Bob Smith")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("Barrett")
	fmt.Println(fn2, sn2)

}
