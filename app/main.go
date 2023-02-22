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
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	practice.Action()
	fmt.Println(practice.SqrtStrict(2))
	fmt.Println(practice.SqrtStrict(-2))
	practice.ReadAction()
	practice.ImageAction()
}
