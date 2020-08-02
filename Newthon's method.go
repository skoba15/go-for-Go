package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
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

func main() {
	fmt.Println(Sqrt(23))
}