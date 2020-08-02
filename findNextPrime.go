package main

import (
	"fmt"
	"math"
	"strconv"
	"bufio"
	"os"
	"strings"
)


func isPrime(x int) bool {
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x % i == 0 {
			return false
		}
	}
	return true
}


func findNextPrimeHelper(x int) {
	if isPrime(x) {
		fmt.Println("Omg next prime found, time to panic!")
		panic(fmt.Sprintf("%v", x))
	}
	defer fmt.Printf("Deferring... %d is not prime\n", x)
	findNextPrimeHelper(x + 1)
}


func findNextPrime(x int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("found prime", r)
		} 
	}()
	findNextPrimeHelper(x + 1)
}



func main() {
	for true {
	    reader := bufio.NewReader(os.Stdin)
	    fmt.Print("Enter number: ")
		Number, _ := reader.ReadString('\n')
		Number = strings.TrimSuffix(Number, "\r\n")
	    num, err := strconv.Atoi(Number)
	    if err != nil {
	    	fmt.Println("Invalid input")
	    } else {
	    	findNextPrime(num)
	    	break
	    }
	}
}