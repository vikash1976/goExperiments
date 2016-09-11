package main

import "fmt"
import "strconv"
//Animal ...
type Animal interface {
	Speak() string
	MutateWeight() 
}
//Cow ...
type Cow struct {
weight float64
}

//Speak ...
func (c Cow) Speak() string {
return `Moo! new weight is: ` + strconv.FormatFloat(c.weight, 'f', -1, 64)
}
//MutateWeight ...
func (c Cow) MutateWeight() {
c.weight = c.weight + 5
fmt.Println("Cow's new weight is: ", c.weight) 
}
//Pig ...
type Pig struct {
weight float64
}
//Speak ...
func (p Pig) Speak() string {
return `Oink! new weight is: `+ strconv.FormatFloat(p.weight, 'f', -1, 64)
}
//MutateWeight ...
func (p *Pig) MutateWeight() {
p.weight = p.weight + 10
fmt.Println("Pig's new weight is: ", p.weight) 
}

func main() {
animals := []Animal{&Cow{60}, &Pig{100}}
for _, animal := range animals {
animal.MutateWeight()
fmt.Println(animal.Speak())
}
}
