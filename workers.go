package main


import (
	"fmt"
	"time"
)


func doWork(workerId int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Printf("%d started working on %d\n", workerId, i)
		time.Sleep(time.Second)
		fmt.Printf("%d finished working on %d\n", workerId, i)
		results <- i*i
	}
}



func main() {
  	numJobs := 7
  	numWorkers := 3

  	jobs := make(chan int, numJobs)
  	results := make(chan int, numJobs)


  	/* staring workers */
  	for i := 1; i <= numWorkers; i++ {
  		go doWork(i, jobs, results)
  	}

  	/* preparing work */
  	for i := 1; i <= numJobs; i++ {
  		jobs <- i
  	}

  	time.Sleep(time.Second)
  	/* closing channel is usually done by the sender, ensuring that no
  	info will be sent again */
  	close(jobs)


  	/* ensuting that all the jobs are done, can be replaced with Waitsyncgroup */
  	for i := 1; i <= numJobs; i++ {
  		<-results
  	}

}