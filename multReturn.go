package main

import "fmt"

//input    //output
func dual(a, b int) (int, bool) {
	return a + b, a > b
}
func vals() (int, int) {
	return 3, 6
}
func multiReturn() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c := vals() //discarding 1 st value by using _ identifier
	fmt.Println(c)
	v1, v2 := dual(10, 20)
	fmt.Println(v1)
	fmt.Println(v2)

}
