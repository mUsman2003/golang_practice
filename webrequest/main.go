package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("WebRequest Practice")
	response, err := http.Get("https://github.com/mUsman2003")
	if err != nil {
		panic(err)
	}
	databyte, err := ioutil.ReadAll(response.Body)
	fmt.Println("Data%v", string(databyte))
	defer response.Body.Close()
}
