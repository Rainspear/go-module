package main

import (
	"fmt"

	"example.com/greetings"
	"example.com/hello"
)

func main() {
	// Get a greeting message from Greeting Module and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	// Say hello from a Hello Module
	fmt.Println(hello.Hello())
}
