package customer

/****
Package with all customer related functions
****/
import (
	"encoding/json"
	"fmt"
)

//Customer type definition
type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

//Customers slice of Customer
type Customers []Customer

//Dummy customers slice
var customers = Customers{
	{
		Name:  "C1",
		Email: "c1@in.com",
		Phone: "9999999999",
	},
	{
		Name:  "C2",
		Email: "c2@in.com",
		Phone: "8899999988",
	},
}

//GetCustomer based on supplied index
func GetCustomer(index int) []byte {
	var js = []byte("{}")
	//to handle any panic
	defer func() {

		if r := recover(); r != nil {
			fmt.Printf("Panic Handled locally: %s\n", r)

		}

	}()

	custToReturn := customers[index-1]
	js, err := json.Marshal(custToReturn)
	if err != nil {
		fmt.Println(err)
		return []byte("Error marshalling customer")
	}
	return js
}

//GetCustomers gets all customers
func GetCustomers() string {
	fmt.Printf("Returning Customer %s\n", customers)
	js, err := json.Marshal(customers)
	if err != nil {
		return "Error marshalling customer"
	}
	return string(js)
}

//UpdateCustomer updates supplied customer, adds for now to the customers slice
func UpdateCustomer(customer []byte) string {
	var c Customer
	err := json.Unmarshal(customer, &c)
	customers = append(customers, c)
	if err != nil {

		return "Error unmarshalling customer"
	}
	js, err := json.Marshal(c)
	if err != nil {

		return "Error marshalling customer"
	}
	return string(js)
}

//DeleteCustomer deletes customer at given index
func DeleteCustomer(i int) string {
	customers = append(customers[:i], customers[i+1:]...)
	return `{message: "Customer deleted"}`
}
