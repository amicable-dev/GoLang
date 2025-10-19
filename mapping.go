package main

import (
	"fmt"
	"maps"
)

func mapping() {
	m := make(map[string]int)

	m["val1"] = 12
	m["val2"] = 15

	fmt.Println("map", m)

	mv1 := m["val1"]
	mv2 := m["val3"]
	fmt.Println("do they exist", mv1, mv2)

	fmt.Println("len: ", len(m))

	delete(m, "val2")
	fmt.Println("updated:", m)
	clear(m)
	fmt.Println("empty:", m)

	_, prs := m["val1"] //_ says that ignore the value and focus on the key if it exist so they give either true or false
	fmt.Println("exist?:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n==n2")
	} else {
		fmt.Println("n!=n2")
	}

}
