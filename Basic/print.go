package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var newvar string
	fmt.Println(newvar)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("write somthing")
	input, _ := reader.ReadString('\n')
	fmt.Println("input", input)
}
