package practice

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	var cuurentIndex int = 0
	var nextIndex int = 1
	var cuurentResultIndex int
	var nextTmp int
	return func() int {
		cuurentResultIndex = cuurentIndex
		nextTmp = cuurentIndex + nextIndex
		cuurentIndex = nextIndex
		nextIndex = nextTmp
		return cuurentResultIndex
	}
}

// func main() {
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }
