package main

import (
	"fmt"
	"time"
)

func switchCases() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Its a Weekend ")
	default:
		fmt.Println("Its a WeekDay")
	}
	//used as if else since we didnt gave any expression with switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("not past 12")
	default:
		fmt.Println("Sleep time")
	}

	WhoAmi := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Im a int")

		case bool:
			fmt.Println("Im a bool")
		default:
			fmt.Println("Dont know maybe ", t)
		}
	}
	WhoAmi("apple")
	WhoAmi(24)
	WhoAmi(false)
}
