package main

import (
	"fmt"
	"sync"
	"time"
)

func task(number int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Task %d started\n", number)

	time.Sleep(1 * time.Second)

	fmt.Printf("Task %d finished\n", number)
}

func main() {
	startTime := time.Now()

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()

	elapsed := time.Since(startTime)

	fmt.Printf("Elapsed time: %.2f seconds\n", elapsed.Seconds())
}
