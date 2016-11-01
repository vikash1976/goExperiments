package main

import (
	"encoding/json"
	"fmt"
	"github.com/vikash1976/goExperiments/go-web/go-http-mongo-gorilla-negroni/customer"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"io"
	
)

var (
    server   *httptest.Server
    reader   io.Reader 
    customersUrl string
	customerUrl string
)

func init() {
    server = httptest.NewServer(Handlers()) //Creating new server with the user handlers

    customersUrl = fmt.Sprintf("%s/api/customers", server.URL) //Grab the address for the API endpoint
	customerUrl = fmt.Sprintf("%s/api/customer/C7", server.URL)
}

func TestCustomersHandler(t *testing.T) {
    
    request, err := http.NewRequest("GET", customersUrl, nil) //Create request with JSON body
	request.Header.Set("x-Auth", "Vikash")

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }

    if res.StatusCode != 200 {
        t.Errorf("Success expected: %d", res.StatusCode) //this means our test failed
    }
}
func TestCustomerHandler(t *testing.T) {
    
    request, err := http.NewRequest("GET", customerUrl, nil) //Create request with JSON body
	request.Header.Set("x-Auth", "Vikash")

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }

    var c customer.Customer
	exp := "c7@in.com"
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &c)
	act := c.Email
	if exp != act {
		t.Fatalf("Expected %s received %s", exp, act)
	}
}

func TestCustomerUpdateHandler(t *testing.T) {
    customerJSON := `{"Name": "C7", "Email": "c7@in.com", "Phone": "7777778877"}`

    reader = strings.NewReader(customerJSON) //Convert string to reader

    request, err := http.NewRequest("PUT", customerUrl, reader) //Create request with JSON body
	request.Header.Set("x-Auth", "Vikash")

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }

   if res.StatusCode != 201 {
        t.Errorf("Success expected: %d", res.StatusCode) //this means our test failed
    }
}

func TestCustomerAddHandler(t *testing.T) {
    customerJSON := `{"Name": "C7", "Email": "c7@in.com", "Phone": "7777778877"}`

    reader = strings.NewReader(customerJSON) //Convert string to reader

    request, err := http.NewRequest("POST", customerUrl, reader) //Create request with JSON body
	request.Header.Set("x-Auth", "Vikash")

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }

   if res.StatusCode != 201 {
        t.Errorf("Success expected: %d", res.StatusCode) //this means our test failed
    }
}

func TestCustomerDeleteHandler(t *testing.T) {
    
    request, err := http.NewRequest("DELETE", customerUrl, nil) //Create request with JSON body
	request.Header.Set("x-Auth", "Vikash")

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }

    if res.StatusCode != 204 {
        t.Errorf("Success expected: %d", res.StatusCode) //this means our test failed
    }
}