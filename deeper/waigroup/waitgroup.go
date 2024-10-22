package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// creating channels
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 3 times starting worker function
	go worker(1, jobs, results)

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}
