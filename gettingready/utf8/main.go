package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %X \t %#x \t %q \n", i, i, i, i, i)
	}
}
