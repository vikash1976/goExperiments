## Synopsis

An  http server using gorilla mux, negroni and mongodb as backend, show casing starter project to implement microservices based on listed features and tools in Go.



## Motivation

A go http server starter project with 
- negroni middleware to implement few routes going through middleware and few not
- integration with MongoDB
- Usage of gorilla mux
- Usage of httptest


## Installation

- download or fork the repository, unzip it, go to go-http-server and do follwoing commands

``` go get gopkg.in/mgo.v2 ```

``` go get github.com/gorilla/mux ```

``` go get github.com/codegangsta/negroni ```

``` go build page/page.go ```

``` go build customer/customer.go ```

``` go run main.go ```

## Tests

``` go test -v ```
