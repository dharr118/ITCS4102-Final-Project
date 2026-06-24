package main

import (
    "fmt"
    "sync"
    "time"
)

func simulatedTask(taskID int, waitGroup *sync.WaitGroup) {
    defer waitGroup.Done()
    time.Sleep(1 * time.Second)
    fmt.Println("Task", taskID, "finished")
}

func main() {
    start := time.Now()
    var waitGroup sync.WaitGroup

    for i := 1; i <= 5; i++ {
        waitGroup.Add(1)
        go simulatedTask(i, &waitGroup)
    }

    waitGroup.Wait()
    fmt.Println("Elapsed:", time.Since(start))
}
