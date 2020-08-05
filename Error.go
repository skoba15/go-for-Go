package main

import (
	"fmt"
	"math"
)


func SqrtHelper(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		change := (z*z - x) / (2*z)
		if math.Abs(change) < 0.001 {
			fmt.Printf("process ended after %d iterations \n", i)
			return z
		}
		z -= change
	}
	return z
}


type ErNegativeSqrt float64


func (e ErNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: -2"
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErNegativeSqrt(x)
	}
	return SqrtHelper(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}