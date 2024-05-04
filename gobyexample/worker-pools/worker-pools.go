package main

import (
	"fmt"
	"time"
)

// worker run several concurrent instances
// Worker receive work on the jobs channal and send corresponding results on `results`
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "started job", j)
		result <- j * 2
	}
}

func main() {
	const numJobs = 5
	// send them work
	jobs := make(chan int, numJobs)
	// collect their results
	results := make(chan int, numJobs)

	// starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs and then close that channel to
	// indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally collect all the resulfts.
	// This also ensures that worker goroutines have finished.
	// An alternative way to wait for multiple goroutines is to use a WaitGroup
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
