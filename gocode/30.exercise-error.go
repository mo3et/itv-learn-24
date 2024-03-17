package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d", float64(e))
}

/*
	func Sqrt(x float64) (float64, error) {
	    return 0, nil
	}
*/
func Sqrt1(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for {
		if math.Abs(z-(z-(z*z-x)/(z*2))) < 0.00000000000001 {
			return z, nil
		} else {
			z = z - (z*z-x)/(z*2)
		}
	}
}

func MySqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, json.Unmarshal([]byte(""), 1)
	}
	z := 1.0
	for i := 0; i < 20; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z, nil
}

func main() {
	z := ErrNegativeSqrt(2).Error()

	fmt.Println(z)
	fmt.Println(Sqrt1(2))
	fmt.Println(Sqrt1(-2))

	if s, err := MySqrt(-2); err != nil {
		fmt.Println(s)
		fmt.Println(err)
	}
}
