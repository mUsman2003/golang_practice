package main

import (
	"fmt"
)

func main() {
	// Call the printMessage function with a sample message
	printMessage("1234")
	var err error
	if 1 == 1 {
		err.errors.new("Error")
	}
}

// Function to print a message
func printMessage(message string) {
	fmt.Println(message)
}
