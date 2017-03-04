package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	//c <- 3
	//c <- 4
	go func() {
		hOt := rand.Intn(2)
		fmt.Println("hOt 11: ", hOt)
		if hOt == 0 {
			c <- 55
		}
	}()
	go func() {
		hOt := rand.Intn(3)
		fmt.Println("hOt 22: ", hOt)
		if hOt == 0 {
			c <- 5
		}
	}()
	go func() {
		/*for d := range c {
			fmt.Println("Data: ", d)
		}*/
		select {
			case d := <-c:
			fmt.Println("Data: ", d)
			fmt.Println("Data1: ", <- c)
			//fmt.Println("Data2: ", <- c)

		}

		fmt.Println("I am done")
	}()
	time.Sleep(20 * time.Second)
	close(c)
	fmt.Println("I am here")
}