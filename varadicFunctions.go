package main

import "fmt"

func sumOFall(nums ...int) {

	fmt.Print("nums :", " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func varFun() {
	sumOFall(1, 3)
	sumOFall(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sumOFall(nums...)
}
