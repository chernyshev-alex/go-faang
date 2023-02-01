package dave

import (
	"fmt"
	"sync"
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

func run_pool() {
	const njobs = 5
	jobs := make(chan int, njobs)
	results := make(chan int, njobs)
	for i := 0; i <= 3; i++ {
		go worker(i, jobs, results)

	}
	for j := 0; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= njobs; a++ {
		<-results
	}
}

func worker_gorgrp(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func wait_grp(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker_gorgrp(i)
		}()

	}
}
