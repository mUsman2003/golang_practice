package main

import "fmt"

func main() {
	fmt.Println("starting pointertsS")

	var ptr *int
	mynum := 100
	fmt.Println("value of num", mynum)
	ptr = &mynum
	fmt.Println("value of ptr", ptr)
	*ptr = *ptr * 10
	fmt.Println("value of ptr", *ptr)
}
