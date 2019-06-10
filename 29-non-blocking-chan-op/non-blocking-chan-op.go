// Basic sends and receives on channels are blocking.
// However, we can use `select` with a `default` clause to
// implement _non-blocking_ sends, receives, and even
// non-blocking multi-way `select`s.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.
	// go func() {
	// 	messages <- "xx"
	// }()

	// select {
	// case msg := <-messages:
	// 	fmt.Println("received message", msg)
	// default:
	// 	fmt.Println("no message received")
	// }

	// A non-blocking send works similarly. Here `msg`
	// cannot be sent to the `messages` channel, because
	// the channel has no buffer and there is no receiver.
	// Therefore the `default` case is selected.
	msg := "hi"

	go func() {
		for {
			select {
			case messages <- msg:
				fmt.Println("sent message", msg)
			default:
				fmt.Println("no message sent")
			}
		}
	}()

	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)

	// if we use default case,  main goroutine will match the default case frist 
	// (faster than goroutine of select messages<-msg), then exit program
	default:
		fmt.Println("no activity")
	}
}
