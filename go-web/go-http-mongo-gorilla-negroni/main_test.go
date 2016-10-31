package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/vikash1976/goExperiments/go-web/go-http-server/customer"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//Testing a GET method
func Test_customerHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8070/customer", nil)
	if err != nil {
		t.Fatal(err)
	}

	var params httprouter.Params
	params = append(params, httprouter.Param{Key: "id", Value: "2"})

	res := httptest.NewRecorder()
	customerHandler(res, req, params)

	//exp := `{"name":"C1","email":"c1@in.com","phone":"9999999999"}`
	//act := res.Body.String()
	var c customer.Customer
	exp := "c2@in.com"
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &c)
	act := c.Email
	if exp != act {
		t.Fatalf("Expected %s received %s", exp, act)
	}
}

//Testing a PUT method
func Test_customerUpdateHandler(t *testing.T) {
	req, err := http.NewRequest("PUT", "http://localhost:8070/customer", strings.NewReader(`{"name":"C3","email":"c3@in.com","phone":"8899997777"}`))
	if err != nil {
		t.Fatal(err)
	}

	var params httprouter.Params
	params = append(params, httprouter.Param{Key: "id", Value: "3"})

	res := httptest.NewRecorder()
	customerUpdateHandler(res, req, params)

	var c customer.Customer
	exp := "c3@in.com"
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &c)
	act := c.Email
	if exp != act {
		t.Fatalf("Expected %s received %s", exp, act)
	}
}

//Testing a DELETE method
func Test_customerDeleteHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8070/customer", nil)
	if err != nil {
		t.Fatal(err)
	}

	var params httprouter.Params
	params = append(params, httprouter.Param{Key: "id", Value: "1"})

	res := httptest.NewRecorder()
	customerDeleteHandler(res, req, params)

	exp := `{message: "Customer deleted"}`
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s received %s", exp, act)
	}
}
