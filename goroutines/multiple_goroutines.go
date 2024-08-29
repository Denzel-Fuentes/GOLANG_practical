package main

import (
    "fmt"
    "time"
)

func worker(id int, done chan bool) {
    fmt.Printf("Worker %d starting\n", id) // Print the start message with the worker ID.
    time.Sleep(time.Second)                // Simulate work by sleeping for one second.
    fmt.Printf("Worker %d done\n", id)     // Print the completion message with the worker ID.
    done <- true                           // Send a completion signal (true) to the 'done' channel.
}


func main() {
    numWorkers := 3 // Number of workers to spawn.
    done := make(chan bool, numWorkers) // Create a buffered channel with a capacity equal to the number of workers.

    for i := 0; i < numWorkers; i++ {
        go worker(i, done) // Start each worker as a goroutine.
    }
    
    for i := 0; i < numWorkers; i++ {
        <-done // Wait for each worker to signal completion before proceeding.
    }
}

