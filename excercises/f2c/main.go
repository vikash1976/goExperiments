//Farenhite to Celcius

package main

import "fmt"

func divisible(i int) {
	if i%2 == 0 {
		fmt.Printf("%v - divisble by 2\n", i)
	} else if i%3 == 0 {
		fmt.Printf("%v - divisble by 3\n", i)
	} else if i%4 == 0 {
		fmt.Printf("%v - divisble by 4\n", i)
	}

	//The conditions are checked top down and the first one to result in true will have its associated block executed. None of the other blocks will execute, even if their conditions also pass. (So for example the number 8 is divisible by both 4 and 2, but the // divisible by 4 block will never execute because the // divisible by 2 block is done first)
}

func fizBuzz() {
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Printf("%d is - FizBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d is - Fiz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d is - Buzz\n", i)
		}
		/*else if i % 15 == 0 {
		    fmt.Printf("%d is - FizBuzz\n", i)
		} - this will never have FizBuzz, because this wil never execute*/

	}
}
func main() {
	var farValue float64

	fmt.Print("Enter a farenhite value: ")
	fmt.Scanf("%f", &farValue)

	celValue := (farValue - 32) * 5 / 9
	fmt.Printf("%f farenhit = %f celcius\n", farValue, celValue)

	for i := 1; i <= 10; i++ {
		divisible(i)
	}

	fizBuzz()

	
}
