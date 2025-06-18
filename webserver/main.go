package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Local host")
	myGetrequest()
}

func myGetrequest() {
	const myurl = "http://localhost:3000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("COntent Length: ", response.ContentLength)
	data, _ := ioutil.ReadAll((response.Body))
	fmt.Println(string(data))
}
