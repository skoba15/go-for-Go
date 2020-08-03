package main

import (
	"fmt"
)


/* example of a closure function. helper returns a closure function. 
returned func has a reference to a and b, so calling that func n times will
actually return nth Fibonacci number
*/


func helper() func() int {
	a := 1
	b := 1
	return func() int {
		result := a
		tmp := b
		b = a + b
		a = tmp
		return result
	}
}


func getFibonacciNumber(number int) int {
	nextFibonacciNumFunc := helper()
	result := 0
	for i := 0; i < number; i++ {
		result = nextFibonacciNumFunc()
	}
	return result
}


func main() {
    fmt.Println(getFibonacciNumber(4))
}