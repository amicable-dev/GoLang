package main

import (
	"fmt"
)

func addition(a, b, c int) int {
	return a + b + c
}

func sol() {
	temp := addition(10, 20, 30)
	fmt.Println(temp)
}

//function can have two parameter ends
//where first () takes input and () decides the out put example
//func example(a,b int)(int,bhool){
//return a+b,a>b
//	}
