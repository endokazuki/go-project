package main

import (
	"go-project/app/practice"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func main() {
	practice.Sqrt(4)
	pic.Show(practice.Pic)
	wc.Test(practice.WordCount)
}
