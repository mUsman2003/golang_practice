package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")
	present := time.Now()
	fmt.Println(present)
	fmt.Println(present.Format("01-02-2006"))
}
