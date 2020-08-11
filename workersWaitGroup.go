package main


import (
	"fmt"
	"time"
	"sync"
)


func doWork(workerId int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for i := range jobs {
		fmt.Printf("%d started working on %d\n", workerId, i)
		time.Sleep(time.Second)
		fmt.Printf("%d finished working on %d\n", workerId, i)
		results <- i*i
		wg.Done()
	}
}



func main() {
  	numJobs := 7
  	numWorkers := 3


  	var wg sync.WaitGroup

  	jobs := make(chan int, numJobs)
  	results := make(chan int, numJobs)


  	/* staring workers */
  	for i := 1; i <= numWorkers; i++ {
  		go doWork(i, jobs, results, &wg)
  	}

  	/* preparing work */
  	for i := 1; i <= numJobs; i++ {
  		wg.Add(1)
  		jobs <- i
  	}

  	time.Sleep(time.Second)
  	/* closing channel is usually done by the sender, ensuring that no
  	info will be sent again */
  	close(jobs)


  	wg.Wait()

}