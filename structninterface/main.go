package main

import (
	"fmt"
	"math"
)
// Shape interface with area method set ...
type Shape interface {
	area() float64
}

// Circle ...
type Circle struct {
	x, y, r float64
}
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Rectangle ...
type Rectangle struct {
	x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// ShapeCollection ...
type ShapeCollection struct {
	shapes []Shape
}
func (m *ShapeCollection) area() float64 {
	var area float64
	for _, s := range m.shapes {
		tempArea := s.area()
		area += tempArea
		fmt.Printf("Area invoked for: %v and area is: %v\n", s, tempArea)
	}

	return area
}
func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 5, 5}

	c1 := Circle{0, 0, 8}
	r1 := Rectangle{0, 0, 15, 15}

	c2 := Circle{0, 0, 3}
	r2 := Rectangle{0, 0, 12, 12}

	m1 := new(ShapeCollection)    // get *MultiShape, hence for m we don't need &m1
	m1.shapes = []Shape{&c1, &r1} // to its shapes field assign value of Shapes array having a circle and a rect
	fmt.Printf("M1's Area: %v\n", m1.area())
	fmt.Println("--------------")
	// Assign ShapeCollection shapes field with array of Shapes with 1 circle and 1 rect - create 1st and then set shapes field
	//m2 := ShapeCollection{}
	//m2.shapes = []Shape{&c2, &r2}
	m2 := ShapeCollection{[]Shape{&c2, &r2}} // initailize

	fmt.Printf("M2's Area: %v\n", m2.area())
	fmt.Println("--------------")
	//fmt.Println(m2.shapes[1].area())
	m := ShapeCollection{[]Shape{&c, &r, m1, &m2}} // initializing m with 1 circle, 1 rect and 2 ShapeCollections

	fmt.Printf("M's Area: %v\n", m.area())
}
