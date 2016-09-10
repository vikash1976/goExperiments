// Go supports _methods_ defined on struct types.

package main

import "fmt"

type rect struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
// This can mutate the receiving struct
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
// This can't mutate the receiving struct
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
// This can't mutate the receiving struct
func (r rect) edit() {
	r.width += 2
	r.height += 5
	fmt.Println("in Edit", r)
}

// This `area` method has a _receiver type_ of `*rect`.
// This can mutate the receiving struct
func (r *rect) editMe() {
	r.width += 2
	r.height += 2
	fmt.Println("In EditMe", r)
}
func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	rp1 := rect{width: 10, height: 5}
	fmt.Println("perim:", rp1.perim())
	rp1.editMe() //Mutated
	fmt.Println("RP1", rp1)
	fmt.Println("perim:", rp1.perim())
	rp2 := &rp1
	rp2.edit() // copy got mutated, not the receiving struct
	fmt.Println("RP2", rp2)
	fmt.Println("perim:", rp2.perim())

	/* /c/Eee/GoWorkspace/src/github.com/vikash1976/goExperiments/structninterface/recievernvalparam(master) $ go run main.go
	area:  50
	perim: 30
	area:  50
	perim: 30
	perim: 30
	In EditMe &{12 7}
	RP1 {12 7}
	perim: 38
	in Edit {14 12}
	RP2 &{12 7}
	perim: 38
	*/

}
