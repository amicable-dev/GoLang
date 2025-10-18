package main

//there is no terinary operator in go so we have to use if/else for even basic conditions

import "fmt"

func conditional() {
	if 8%3 == 1 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	if num := 3; num > 4 {
		fmt.Println("True")
	} else if num > 1 && num < 5 {
		fmt.Println("dual cond")
	}

}
