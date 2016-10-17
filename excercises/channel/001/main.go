package main

import (
	"fmt"
    "time"
)

func main() {
    //making the channel, the unbuffered one
	c := make(chan int, 5)

	//Go launches this annonymous function as a seperate process
    go func() {
		for i := 0; i < 10; i++ {	
            fmt.Printf("Put %d onto the channel\n", i)		
            c <- i            			
		}
		close(c) //closing the channel to indicate main that no more writes from here on it.
	}()

//close(c)
	//range c - making main to wait for another process running our
    //annonymous function to write to channel c. The wait loop will only
    //get over when the channel is closed, meaning nothing more to read from it.
    time.Sleep(1 * time.Second)
    for n := range c {
		fmt.Println(n)
	}
}