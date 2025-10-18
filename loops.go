package main

import "fmt"

func loops() {

	for j := 1; j < 4; j++ {
		fmt.Println(j)
	}

	for i := range 5 {
		fmt.Println("range = ", i)
	}

	i := 1
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
