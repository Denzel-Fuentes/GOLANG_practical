package main

/* Working with threads in Go, commonly called goroutines, 
is one of the fundamental features of the language. Go handles
concurrency in a way that differs significantly from other
programming languages ​​such as Java or C++. Goroutines are 
much lighter than operating system threads and are handled internally by the Go runtime scheduler.
 */
import(
	"fmt"

)

func sayHello()  {
	fmt.Println("Hello")
}

//Synchonization with Channels
func worker(done chan bool)  {
	fmt.Print("Working...")
	fmt.Println("Done")
	done <- true
}

func main()  {
    // Starting a Goroutine
    go sayHello()
    
    // Goroutine with Anonymous Functions
    go func() {
        fmt.Println("Hello from goroutine!!")
    }()
    
    // Creating a channel 'done' that can hold one boolean value
    done := make(chan bool, 1)
    
    // Starting the 'worker' function as a Goroutine and passing the 'done' channel to it
    go worker(done)
    
    // Blocking the main Goroutine until a signal (boolean true) is received from the 'worker' Goroutine
    <-done
    fmt.Println("End program")
}
