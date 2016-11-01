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
)

//Customer type definition
type Customer struct {
	ID string `bson:"ID"`
	Name  string `bson:"Name"`
	Email string `bson:"Email"`
	Phone string `bson:"Phone"`
	Category string `bson:"Category"`
	Networth float64 `bson:"Networth"`
}

//Customers slice of Customer
type Customers []Customer


//GetCustomer returns customer based on provided index
func GetCustomer(s *mgo.Session, custID string) ([]byte, error) {
	session := s.Copy()
	defer session.Close()
	var customer1 Customer
	collection := session.DB("customers").C("customers")

	err := collection.Find(bson.M{"ID": custID}).One(&customer1)

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

	err = collection.Update(bson.M{"ID": c.ID}, &c)
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
		Key:        []string{"ID"},
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
//CategoryTotal category wise totla networth
type CategoryTotal struct {
	Category string
	Amount   float64
}
//GetCustomerCategoryTotals returns total on networth category wise
func GetCustomerCategoryTotals(s *mgo.Session) ([]byte, error) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("customers").C("customers")
	
	var categoryTotals []CategoryTotal

	//Group by category, the amount spent on a category
	pipe := c.Pipe([]bson.M{{"$group": bson.M{"_id": "$Category",
		"TotalAmount": bson.M{"$sum": "$Networth"}}}})
	iter := pipe.Iter()
	var x map[string]interface{}
	for iter.Next(&x) {
		categoryTotals = append(categoryTotals,
			CategoryTotal{Category: x["_id"].(string),
				Amount: x["TotalAmount"].(float64)})
	}
	js, err := json.Marshal(categoryTotals)
	if err != nil {
		Err := fmt.Errorf("Error marshaling: %v", err)
		return nil, Err
	}
	return js, nil
	
}
