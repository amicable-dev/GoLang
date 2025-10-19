package main

import (
	"fmt"
)

func mapping() {
	m := make(map[string]int)

	m["val1"] = 12
	m["val2"] = 15

	fmt.Println("map", m)

	mv1 := m["val1"]
	mv2 := m["val3"]
	fmt.Println("do they exist", mv1, mv2)
}
