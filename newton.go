package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

func Sqrt(x float64) (float64, *ErrNegativeSqrt) {
	if x < 0 {
		error := ErrNegativeSqrt(x)
		return 1.0, &error
	}
	z := 1.0
	zChange := 1.0
	for math.Abs(zChange) > 0.00000000000001 {
		fmt.Println(z)
		zChange = (z*z - x) / (2*z)
		z -= zChange
	}
	return z, nil
}

