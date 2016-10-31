package customer

/****
Package with all customer related functions
****/
import (
	"encoding/json"
	"log"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

//Customer type definition
type Customer struct {
	Name  string `bson:"Name"`
	Email string `bson:"Email"`
	Phone string `bson:"Phone"`
}


func connect() (session *mgo.Session) {
	connectURL := "localhost"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	return session
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


//GetCustomer returns customer based on provided index
func GetCustomer(custID string) []byte {
	session := connect()
	defer session.Close()
	var customer1 Customer
	collection := session.DB("customers").C("customers")

	err := collection.Find(bson.M{"Name": custID}).One(&customer1) 
	
	if err != nil {
            
            log.Println("Failed get customer: ", err)
            return []byte("Customer Not Found")
    }
	log.Printf("DB Results  %v\n", customer1)
	respBody, err := json.Marshal(customer1)
	log.Printf("Returning Customer %s\n", customer1)
	return respBody
}

//GetCustomers gets all customers
func GetCustomers() string {
	
	session := connect()
	defer session.Close()
	var customers1 []Customer
	collection := session.DB("customers").C("customers")

	err := collection.Find(bson.M{}).All(&customers1) 
	
	if err != nil {
            
            log.Println("Failed get all customers: ", err)
            return "Database error"
    }
	log.Printf("DB Results  %v\n", customers1)
	respBody, err := json.Marshal(customers1)
	log.Printf("Returning Customer %s\n", customers1)
	return string(respBody)
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
