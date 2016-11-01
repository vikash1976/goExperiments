package customer

/****
Package with all customer related functions
****/
import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	//"os"
	//"errors"
)

//Customer type definition
type Customer struct {
	Name  string `bson:"Name"`
	Email string `bson:"Email"`
	Phone string `bson:"Phone"`
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
func GetCustomer(s *mgo.Session, custID string) ([]byte, error) {
	session := s.Copy()
	defer session.Close()
	var customer1 Customer
	collection := session.DB("customers").C("customers")

	err := collection.Find(bson.M{"Name": custID}).One(&customer1)

	if err != nil {

		log.Println("Failed get customer: ", err)
		Err := fmt.Errorf("Failed to get customer: %v", err)
		return nil, Err
	}
	log.Printf("DB Results  %v\n", customer1)
	respBody, err := json.Marshal(customer1)
	log.Printf("Returning Customer %s\n", customer1)
	return respBody, nil
}

//GetCustomers gets all customers
func GetCustomers(s *mgo.Session) ([]byte, error) {

	session := s.Copy()
	defer session.Close()
	var customers1 []Customer
	collection := session.DB("customers").C("customers")

	err := collection.Find(bson.M{}).All(&customers1)

	if err != nil {

		log.Println("Failed get customers: ", err)
		Err := fmt.Errorf("Failed to get customers: %v", err)
		return nil, Err
	}
	log.Printf("DB Results  %v\n", customers1)
	respBody, err := json.Marshal(customers1)
	log.Printf("Returning Customer %s\n", customers1)
	return respBody, nil
}

//AddCustomer adds supplied customer
func AddCustomer(s *mgo.Session, customer []byte) ([]byte, error) {
	session := s.Copy()
	defer session.Close()
	var c Customer
	err := json.Unmarshal(customer, &c)

	collection := session.DB("customers").C("customers")

	err = collection.Insert(c)
	if err != nil {
		if mgo.IsDup(err) {
			log.Println("Customer with this Name already exists: ", err)
			Err := fmt.Errorf("Customer with this Name already exists: %v", err)
			return nil, Err

		}

		log.Println("Database Error: ", err)
		Err := fmt.Errorf("Database Error: %v", err)
		return nil, Err
	}
	js, err := json.Marshal(c)
	if err != nil {
		Err := fmt.Errorf("Error marshaling: %v", err)
		return nil, Err
	}
	return js, nil
}

//UpdateCustomer updates supplied customer
func UpdateCustomer(s *mgo.Session, customer []byte) ([]byte, error) {
	session := s.Copy()
	defer session.Close()
	var c Customer
	err := json.Unmarshal(customer, &c)

	collection := session.DB("customers").C("customers")

	err = collection.Update(bson.M{"Name": c.Name}, &c)
	if err != nil {
		switch err {
		default:
			log.Println("Failed update customer: ", err)
			Err := fmt.Errorf("Failed update customer: %v", err)
			return nil, Err
		case mgo.ErrNotFound:
			log.Println("Customer not found: ", err)
			Err := fmt.Errorf("Customer not found: %v", err)
			return nil, Err
		}
	}
	js, err := json.Marshal(c)
	if err != nil {
		Err := fmt.Errorf("Error marshaling: %v", err)
		return nil, Err
	}
	return js, nil
}

//EnsureIndex ensures that primary key is set and reports violations accordingly
func EnsureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("customers").C("customers")

	index := mgo.Index{
		Key:        []string{"Name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}
}

//DeleteCustomer deletes customer at given index
func DeleteCustomer(s *mgo.Session, custID string) ([]byte, error) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("customers").C("customers")

	err := c.Remove(bson.M{"Name": custID})
	if err != nil {
		switch err {
		default:
			log.Println("Failed deleting customer: ", err)
			Err := fmt.Errorf("Failed update customer: %v", err)
			return nil, Err
		case mgo.ErrNotFound:
			log.Println("Customer not found: ", err)
			Err := fmt.Errorf("Customer not found: %v", err)
			return nil, Err
		}
	}
	return []byte(`{message: "Customer deleted"}`), nil
}
