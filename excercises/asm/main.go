package main

import (
	"fmt"
)

/*func getSmallest(x []float64) float64{

	var smallest = x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < smallest {
			smallest = x[i]
		}

	}
	return smallest
}*/
func getSmallest(x []float64) (smallest float64) { //named return

	smallest = x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < smallest {
			smallest = x[i]
		}

	}
	return
}

func retMultiple() (int, bool) {
	return 5, true
}

//Variadic
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	var total float64
	x := [5]float64{
		98,
		93,
		77,
		82,
		83, //Notice the extra trailing , after 83. This is required by Go and it allows us to easily remove an element from the array by commenting out the line
	}

	for _, value := range x { //use of blank identifier, we know we get it, but we don't want to use it, range returns current index and value of element
		total += value
	}
	fmt.Printf("Average: %f\n", total/float64(len(x)))

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3) // swap the place and result is 0,0,3 and 0,0
	fmt.Println(slice3, slice4)

	arr := [5]float64{1, 2, 3, 4, 5}
	s := arr[0:5]

	fmt.Println(s)
	series := []float64{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := getSmallest(series)
	fmt.Println(smallest)

	ret, status := retMultiple()

	fmt.Println(ret, status)

	xs := []int{1, 2, 3}
	fmt.Println(add(xs...)) //... after slice sends each element seperately
	fmt.Println(add(1, 2, 3))

}
