package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://reqres.in/api/products/3"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Body)
	data, err := ioutil.ReadAll(response.Body)
	fmt.Println(strings(data))
}
