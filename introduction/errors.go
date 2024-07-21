package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return errors.New("unable to do something")
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success")
	}
}
