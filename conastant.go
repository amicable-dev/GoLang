package main

import (
	"fmt"
	"math"
)

const a string = "hello"

func constant() {
	fmt.Println(a)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
