package main

import "fmt"

func main() {

    // Create a new channel with `make(chan val-type)`.
    // Channels are typed by the values they convey.
    messages := make(chan string)

    // _Send_ a value into a channel using the `channel <-`
    // syntax. Here we send `"ping"`  to the `messages`
    // channel we made above, from a new goroutine.
    go func() {
        messages <- "deep"
        messages <- "pjfjfjjfjf"
    }()
    msg := <-messages
    fmt.Println("1st print:", msg)
    go func() { 
        messages <- "ping" 
        messages <- "p2"
    }()

    go func(){
        messages <- "final"
    }()

    fmt.Println(msg)
    // The `<-channel` syntax _receives_ a value from the
    // channel. Here we'll receive the `"ping"` message
    // we sent above and print it out.
    m2 := <-messages
    
    fmt.Println( m2)

   
}