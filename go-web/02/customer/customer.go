package customer

import (
    "fmt"
    "encoding/json"
)

type Customer struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Phone string `json:"phone"`
}

type Customers []Customer

var customers  = Customers {
       { 
        Name: "C1",
        Email: "c1@in.com",
        Phone: "9999999999",
    },
    {
        Name: "C2",
        Email: "c2@in.com",
        Phone: "8899999988",
    },
}

//GetCustomer based on supplied index
func GetCustomer(index int) []byte {
    var js = []byte("{}")
    defer func() {
        
        if r := recover(); r != nil {
            fmt.Printf("Panic Handled locally: %s\n", r)        
            
        }
        
    }()
    //fmt.Printf("Returning Customer %s\n", customers[index -1 ])
    custToReturn := customers[index - 1]
    js, err := json.Marshal(custToReturn)
    if err != nil {
        fmt.Println(err)
        return []byte("Error marshalling customer")
    }
    fmt.Println("Returning from GetCustomer")
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