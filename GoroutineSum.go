package main


import "fmt"


func sum(numbers []int, start int, n int, channel chan int) {
	result := 0
	for i := start; i < n; i += 2 {
		result += numbers[i]
	}
	channel <- result
}


func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7}

    channel := make(chan int)

    go sum(numbers, 0, len(numbers), channel)
    go sum(numbers, 1, len(numbers), channel)

    fmt.Printf("first written sum is is: %d\n", <-channel)
    fmt.Printf("second written sum is: %d", <-channel)
}