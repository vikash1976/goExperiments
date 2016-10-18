package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
    "strings"
    "io/ioutil"
    "github.com/vikash1976/goExperiments/go-web/02/customer"
)
//Testing a GET method
func Test_customerHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "http://localhost:8070/customer", nil)
    if err != nil {
        t.Fatal(err)
    }
    
    var params httprouter.Params
    params = append(params, httprouter.Param {Key: "id", Value: "2"})
    

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
    req, err := http.NewRequest("PUT", "http://localhost:8070/customer", strings.NewReader(`{"name":"C2","email":"c2@in.com","phone":"8899997777"}`))
    if err != nil {
        t.Fatal(err)
    }
    
    var params httprouter.Params
    params = append(params, httprouter.Param {Key: "id", Value: "3"})
    

    
    res := httptest.NewRecorder()
    customerUpdateHandler(res, req, params)

    //exp := `{"name":"C1","email":"c1@in.com","phone":"9999999999"}`
    //act := res.Body.String()
    var c customer.Customer
    exp := "c4@in.com"
    body, _ := ioutil.ReadAll(res.Body)
    err = json.Unmarshal(body, &c)
    act := c.Email
    if exp != act {
        t.Fatalf("Expected %s received %s", exp, act)
    }
}



