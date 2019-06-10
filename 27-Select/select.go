// Go’s select lets you wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful feature of Go.
package main

import "time"
import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time, to simulate e.g. 
	// blocking RPC operations executing in concurrent goroutines.

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	

	// We’ll use select to await both of these values simultaneously, 
	// printing each one as it arrives.

	for i := 0; i < 2; i++ {
		select {
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg1 := <-c1:
			fmt.Println("receive", msg1)
		}
	}
}
