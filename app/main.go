package main

import (
	"fmt"
	"go-project/app/practice"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func main() {
	practice.Sqrt(4)
	pic.Show(practice.Pic)
	wc.Test(practice.WordCount)
	f := practice.Fibonacci()
	for i := 1; i < 10; i++ {
		fmt.Println(f(i))
	}
}
