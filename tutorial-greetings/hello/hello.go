package main

import (
	"fmt"

	"tutorial.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Eunbin")
	fmt.Println(message)
}