package main

import (
	"fmt"
	"slices"
)

func slice() {
	var str []string
	fmt.Println("uninit: ", str, str == nil, len(str) == 0)

	str = make([]string, 3)
	fmt.Println("emp: ", str, " len ", len(str), " cap ", cap(str))

	str[0] = "a"
	str[1] = "b"
	str[2] = "c"
	fmt.Println("set:", str)      //whole string array
	fmt.Println("get:", str[2])   //should give c
	fmt.Println("len:", len(str)) //3

	str = append(str, "d") //add the value to the last
	str = append(str, "e", "f")
	fmt.Println("append: ", str)

	str2 := make([]string, len(str))
	copy(str2, str)
	fmt.Println("coppy: ", str2)

	l := str[2:5]
	fmt.Println("Slice1: ", l)

	l = str[:5]
	fmt.Println("Slice2: ", l)

	l = str[2:]
	fmt.Println("Slice3: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("T: ", t)

	t2 := []string{"a", "b", "c"}
	if slices.Equal(t, t2) {
		fmt.Println("t==t2")
	} else {
		fmt.Println("t!=t2")
	}

}
