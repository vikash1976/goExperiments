package main

import "fmt"

func toPanic() {
    /*defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic Handled locally: %s\n", r)
		}
	}()*/
    arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr[7])
    fmt.Println("After Panic")
}

//var x int

func main() {
	/*defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic Handled in caller (main): %s\n", r)
		}
	}()*/
	toPanic()
	fmt.Println("After panick handled - normal flow in main")
}
