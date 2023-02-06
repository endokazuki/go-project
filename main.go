package main

import (
	"fmt"

	"github.com/endokazuki/go-project/calculate"
)

func main() {
	// fmt.Println("When's Saturday?")
	// fmt.Println(time.Now().Weekday())
	// today := time.Now().Weekday()
	// switch time.Saturday {
	// case today + 0:
	// 	fmt.Println("Today.")
	// case today + 1:
	// 	fmt.Println("Tomorrow.")
	// case today + 2:
	// 	fmt.Println("In two days.")
	// default:
	// 	fmt.Println("Too far away.")
	// }
	fmt.Println(calculate.Sqrt(4))
}
