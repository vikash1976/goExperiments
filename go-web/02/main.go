package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/vikash1976/goExperiments/go-web/02/page"
	"net/http"
	"regexp"
	"github.com/vikash1976/goExperiments/go-web/02/customer"
	"io"
	"io/ioutil"
	"strconv"
	)

var validPath = regexp.MustCompile("([a-zA-Z0-9]+)$")

//function with closure to bring title check and error handling at one place
func makeHandler(fn func(http.ResponseWriter, *http.Request, httprouter.Params, string)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		m := validPath.FindStringSubmatch(params.ByName("file"))
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, params, m[1])
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, title string) {
	
	p, err := page.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	page.RenderTemplatePage(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, title string) {
	
	p, err := page.LoadPage(title)
	if err != nil {
		p = &page.Page{Title: title}
	}
	page.RenderTemplatePage(w, "edit", p)

}
func saveHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, title string) {
	
	body := r.FormValue("body")
	p := &page.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func customersHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	setHeaders(&w)

	io.WriteString(w, customer.GetCustomers())
	
}
func customerHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	
	custID, _ := strconv.Atoi(params.ByName("id"))
	cust := customer.GetCustomer(custID)
	
	if len(cust) == 0 {
		http.Error(w, "Invalid Index", http.StatusInternalServerError)
		return
	}
	setHeaders(&w)
	io.WriteString(w, cust)
	
}

func customerUpdateHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//fmt.Fprintln(w, "post update")
	fmt.Println("Param is: ", params)
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("some error")
    }
    fmt.Println(string(body))
	 //var t map[string]interface{}
	cust := customer.UpdateCustomer(body)
	setHeaders(&w)
    io.WriteString(w, cust)
	
}
func customerDeleteHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	custID, _ := strconv.Atoi(params.ByName("id"))
	cust := customer.DeleteCustomer(custID)
	setHeaders(&w)
    io.WriteString(w, cust)
	
}

func setHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json; charset=utf-8")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")

}

func main() {
	router := httprouter.New()
	router.GET("/view/:file", makeHandler(viewHandler))
	router.GET("/edit/:file", makeHandler(editHandler))
	router.POST("/save/:file", makeHandler(saveHandler))
	router.GET("/api/customer", customersHandler)
	router.GET("/api/customer/:id", customerHandler)
	router.PUT("/api/customer/:id", customerUpdateHandler)
	router.DELETE("/api/customer/:id", customerDeleteHandler)
	/*http.HandleFunc("/view/",  makeHandler(viewHandler))
	  http.HandleFunc("/edit/",  makeHandler(editHandler))
	  http.HandleFunc("/save/",  makeHandler(saveHandler))*/
	http.ListenAndServe(":8070", router)
	//fmt.Println(customer.GetCustomer(1))
}
