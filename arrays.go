package main

import "fmt"

func arr() {
	//empty array
	var a [5]int
	fmt.Println("Empty : ", a)
	//setting a value
	a[3] = 100
	fmt.Println("SetVal : ", a)
	//Gettin the val
	fmt.Println("SetVal : ", a[3])
	//Finding Array Length
	fmt.Println("len : ", len(a))
	//faster implimentation
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Dcl : ", b)
	var c = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Dcl : ", c)
	//directly assigning 0 val to earlier values
	b = [...]int{100, 200, 4: 400}
	fmt.Println("0 valAssigner : ", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D array : ", twoD)

	twoDAr := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2Darr : ", twoDAr)

}
