package main

import "fmt"

func main() {

	var message = "Hello World!"
	var a1, b1, c1 = 1, false, 3
	var a2, b2 = "A2", "B2"
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'
	g = 'T'
	// := This shorthand is only allowed to be used within functions.

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	fmt.Printf("%v \n", message)
	fmt.Printf("%v \n", a1)
	fmt.Printf("%v \n", b1)
	fmt.Printf("%v \n", c1)

	fmt.Printf("%v \n", a2)
	fmt.Printf("%v \n", b2)

	const (
		a5 = iota
		b5
		c5
		d5
	)
	fmt.Printf("%d\t%d\t%d\t%d\n", a5, b5, c5, d5)

	var x, y float32 = -1, -2 // float32 applies to both x and y
	var (
		i       int
		u, v, s = 2.0, 3.0, "bar"
	)

	fmt.Printf("%v\t %v\t %v\t %v\t %v\t %v \n", x, y, i, u, v, s)
	fmt.Printf("%T\t %T\t %T\t %T\t %T\t %T \n", x, y, i, u, v, s)

	var t int
	fmt.Println(t)
	// value is set to the zero value for its type: false for booleans, 0 for integers, 0.0 for floats, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps. This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified.
	// In addition there two alias types: byte which is the same as uint8 and rune which is the same as int32
	// There are also 3 machine dependent integer types: uint, int and uintptr. They are machine dependent because their size depends on the type of architecture you are using.
	
	//A boolean value (named after George Boole) is a special 1 bit integer type used to represent true and false (or on and off)
}
