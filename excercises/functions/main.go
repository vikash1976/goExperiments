package main

import "fmt"

func nextEven() func() uint { //closure
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint { //recursion
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func firstCall(in int) int {
	fmt.Println("1st Function")
	return in * 2
}
func secondCall() {
    fmt.Println("2nd Function")
}

func calculateInterest(principal float64) func(float64) func(float64) func() float64 { //closure

	return func(time float64) func(float64) func() float64 {
		return func(rate float64) func() float64 {
			return func() float64 {
				return (principal * time * rate) / 100
			}

		}
	}

}

func add2Numbers(op1 int, op2 int) int {
    return op1 + op2 
}
func main() {
	fmt.Println(add2Numbers(5, 4))
	defer func(){
        str := recover()
	    fmt.Printf("Panic Handled: %s\n", str)
    }()
    defer secondCall() //Basically defer moves the call to secondCall to the end of the function, and it will run before the function main returns
    getNextEven := nextEven()
    ret := firstCall(10)
	if ret == 20 {
		return
	}
	for x := 0; x < 10; x++ {
		fmt.Println(getNextEven()) // gets 10 even numbers, i and the func returned from nextEven forms a closure, value of i persists across calls
	}

	fmt.Println(factorial(10))
    
    for100 := calculateInterest(100)
	for100n5yrs := for100(5)
	for100n5yrsAt10 := for100n5yrs(10.25)
	
	fmt.Println("\n", for100n5yrsAt10())
    
    for100n7yrs := for100(7)
	for100n7yrsAt7_75 := for100n7yrs(7.75)
	fmt.Println("\n", for100n7yrsAt7_75())
    
    arr := [] int {1,2,3,4,5}
    fmt.Println(arr[7])
    fmt.Println("After panic handled")

}
