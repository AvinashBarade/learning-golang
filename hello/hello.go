package main

import (
	"fmt"
	"learning-golang/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
