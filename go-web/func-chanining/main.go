package main
import (
    "fmt"
)
type human interface {
    wakeup() human
    getReadytoWork() human
    doWork() human
    packupFortheDay() human
    haveDinnerEtc() human
    gotoBed() human
}
type person struct {
    firstName string
}
func (p person) wakeup() human {
    fmt.Printf("%s woke up\n", p.firstName)
    return p
}
func (p person) getReadytoWork() human {
    fmt.Printf("%s is getting ready to work\n", p.firstName)
    return p
}
func (p person) doWork() human {
    fmt.Printf("%s is cracking the nuts\n", p.firstName)
    return p
}
func (p person) packupFortheDay() human {
    fmt.Printf("%s is packing up for the day\n", p.firstName)
    return p
}
func (p person) haveDinnerEtc() human {
    fmt.Printf("%s is having dinner\n", p.firstName)
    return p
}
func (p person) gotoBed() human {
    fmt.Printf("%s is going to bed, Lets catch up tomorrow morning.\n", p.firstName)
    return p
}
func main() {
    p := person { "John" }
    p.wakeup().getReadytoWork().doWork().packupFortheDay().haveDinnerEtc().gotoBed()
}