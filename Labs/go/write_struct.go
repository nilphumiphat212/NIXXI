package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type TestStruct struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

func main() {
	fmt.Println("Write Struct")
	var data []TestStruct
	for i := 1; i <= 1000; i++ {
		data = append(data, TestStruct{strconv.Itoa(i), "nil", true})
	}
	jsonFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("test.json", jsonFile, 0644)
}
