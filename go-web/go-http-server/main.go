package main

import (
	"log"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net"
	"regexp"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"github.com/codegangsta/negroni"
	"github.com/vikash1976/goExperiments/go-web/go-http-server/customer"
	"github.com/vikash1976/goExperiments/go-web/go-http-server/page"
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
	w.Write(cust) //another way of writing, edited GetCustomer to return []byte instaed of string

}

func customerUpdateHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//log.Fprintln(w, "post update")
	log.Println("Param is: ", params)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("some error")
	}

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

func handleAuthRoutes (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	p := &page.Page{Title: "Deal", Body: []byte("Here is the deal, Enjoi!!!")}
	page.RenderTemplatePage(w, "getToken", p)
}

func myMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  ip, port, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
    }

    userIP := net.ParseIP(ip)
    if userIP == nil {
        fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
        return
    }

    // This will only be defined when site is accessed via non-anonymous proxy
    // and takes precedence over RemoteAddr
    // Header.Get is case-insensitive
    forward := r.Header.Get("X-Forwarded-For")

    log.Printf("Client's IP: %s\n", ip)
    log.Printf("Client's Port: %s\n", port)
    log.Printf("Client's Forwarded for: %s\n", forward)
  next(w, r)//call next function on the stack of middleware
 
}
func main() {
	router := httprouter.New() //instantiating router - httprouter
	//advanced router http://www.gorillatoolkit.org/pkg/mux with more features 
	
	//mapping request paths to handlers
	router.GET("/view/:file", makeHandler(viewHandler))
	router.GET("/edit/:file", makeHandler(editHandler))
	router.POST("/save/:file", makeHandler(saveHandler))
	router.GET("/api/customers", customersHandler)
	router.GET("/api/customer/:id", customerHandler)
	router.PUT("/api/customer/:id", customerUpdateHandler)
	router.DELETE("/api/customer/:id", customerDeleteHandler)
	router.GET("/auth/whatsthedeal", handleAuthRoutes)
	

	//With the middleware we can stack up set of functions that will run for each of
	//our requests before we actually land into handler function
	n := negroni.New() //instatiating a middlewrae - negroni
	n.Use(negroni.NewLogger())//adding Logger middleware
	n.Use(negroni.HandlerFunc(myMiddleware))//adding our custom function as middleware
  	n.UseHandler(router)//configuring middleware to use router for all request to handler mapping
	http.ListenAndServe(":8070", n)

}
