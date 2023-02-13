package practice

import (
	"fmt"
	"math"
)

type errors interface {
	Error() error
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() error {
	if e < 0 {
		err := fmt.Errorf("cannot Sqrt negative number: %v", e)
		return err
	}
	return nil
}

func SqrtStrict(x float64) (float64, error) {
	var Sqrt errors = ErrNegativeSqrt(x)
	return math.Sqrt(x), Sqrt.Error()
}

// func main() {
// 	fmt.Println(SqrtError(2))
// 	fmt.Println(SqrtError(-2))
// }
