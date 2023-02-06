package calculate

var z float64 = 1

func Sqrt(x float64) float64 {
	i := 0
	for i < 10 {
		// fmt.Println(float64(z))
		z -= (float64(z)*float64(z) - float64(x)) / (2 * float64(z))
		i++
	}
	return z
}

// func main() {
// 	fmt.Println(Sqrt(4))
// }
