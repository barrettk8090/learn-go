package greetings

import "fmt"

// Hello returns a greeting for the named person.

func Hello(name string) string {
	// Return a greeting for the named person.
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}
