// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

package main

import (
    "fmt"
    "time"
)

// This is the function we'll run in a goroutine. The `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    // Send a value to notify that worker is done.
    done <- true
}
// This `ping` function only accepts a channel for sending values. It would be a 
//compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
    pings <- msg
}
// The `pong` function accepts one channel for receives (`pings`) and a second for sends (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}
func main() {

	// Here we `make` a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel without 
    // a corresponding concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	//fmt.Println(<-messages)
	// With buffered channel even we don't receive, its fine.
	//If we remove 2 out of make at line 16, we make it unbuffered and without <- chan 
    //program panic with
	/* fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan send]:
	main.main()
	        C:/Eee/GoWorkspace/src/github.com/vikash1976/goExperiments/excercises/channel/main.go:21 +0x75
	exit status 2
	*/
    
    // We can use channels to synchronize execution across goroutines. we will call worker in a
    //goroutine and will synchronize its completion on channel 'done'
    // Start a worker goroutine, giving it 'done' channel to notify on.
    done := make(chan bool, 1)
    go worker(done)

    // Block until we receive a notification from the worker on the channel.
    <-done
    fmt.Println("Worker completed")
    
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message to ping")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
