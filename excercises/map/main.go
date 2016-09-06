package main
import (
	"fmt"
)
func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])

	/*Accessing an element of a map can return two values instead of just one. The first value is
	  the result of the lookup, the second tells us whether or not the lookup was successful.*/
	name, ok := elements["Un"] 
	fmt.Println(name, ok)
	if name, ok := elements["Be"]; ok {
		fmt.Println(name, ok)
	}
    fmt.Println(name, ok)
}




// we can take advantage of it and code like
	//if name, ok := elements["Un"]; ok {
	//	fmt.Println(name, ok)
	//}
	//fmt.Println(name, ok) // undefined name, as the construct above only defines name and of if ok is true