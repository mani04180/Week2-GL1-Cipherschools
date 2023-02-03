package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type something struct {
	Name    string  `json:"name"`
	Married string  `json:"married"`
	Age     float64 `json:"age"`
	Address struct {
		Street int    `json:"street"`
		City   string `json:"city"`
	} `json:"address"`
	Marks []int `json:"marks"`
}

func main() {
	jsonFile, err := os.Open("something.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonByteValues, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var something something
	json.unmarshal(jsonByteValues, &something)
	log.Println(something)
	fmt.Println(jsonByteValues)
}
