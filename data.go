package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One int    `json:"one"`
	two string `json:"two"`
}

func testData1() {
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in)
	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded))
	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out)
}

func testData2() {
	a := []int{1, 2, 3, 4}
	result := make([]*int, len(a))
	for i, v := range a {
		result[i] = &v
	}
	for _, u := range result {
		fmt.Printf("%d ", *u)
	}
}
