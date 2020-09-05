package main

import (
	"fmt";
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	zChange := 1.0
	for math.Abs(zChange) > 0.00000000000001 {
		fmt.Println(z)
		zChange = (z*z - x) / (2*z)
		z -= zChange
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
