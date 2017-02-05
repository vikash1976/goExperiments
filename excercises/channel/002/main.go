package main

import (
    "time"
    "fmt"
)

func main() {
    var Ball int
    table := make(chan int)
    go player(table)
    go player(table)
    go player(table)
    go player(table)

    table <- Ball
    time.Sleep(5 * time.Second)
    fmt.Println(<-table)
}

func player(table chan int) {
    for {
        ball := <-table
        //fmt.Printf("Ball is: %v\n", ball)
        ball++
        time.Sleep(100 * time.Millisecond)
        table <- ball
    }
}