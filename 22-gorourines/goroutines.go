// A _goroutine_ is a lightweight thread of execution.

package main

import "fmt"

func f(from string) {
    for i := 0; i < 80; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    // Suppose we have a function call `f(s)`. Here's how
    // we'd call that in the usual way, running it
    // synchronously.
    

    // To invoke this function in a goroutine, use
    // `go f(s)`. This new goroutine will execute
    // concurrently with the calling one.
	go f("goroutine")
	go f("xxx")
	go f("going")

    // You can also start a goroutine for an anonymous
    // function call.
    // go func(msg string) {
	// 	for i:=0; i < 80; i++{

	// 		fmt.Println(msg)
	// 	}
    // }("going")

    // Our two function calls are running asynchronously in
    // separate goroutines now, so execution falls through
    // to here. This `Scanln` requires we press a key
    // before the program exits.
    fmt.Scanln()
    fmt.Println("done")
}
