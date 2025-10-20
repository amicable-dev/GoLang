package main

import "fmt"

func sumOFall(nums ...int) {

	fmt.Print("nums :", " ")
	total := 0
	for _, num := range nums { //here we using _ because nums... becoms []int slice which comes with key and value so values get stored into v and the index is getting ignored
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
