package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"github.com/vikash1976/goExperiments/go-web/go-http-server-gorilla-mux/page"
	"github.com/vikash1976/goExperiments/go-web/go-http-server-gorilla-mux/customer"
	"github.com/codegangsta/negroni"
	"net"
	
)

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All categories")
	vars := mux.Vars(r)
	category := vars["category"]
	sort := vars["sort"]
	fmt.Printf("Req: %v\n%v\n", category, sort)
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Particular category")
	vars := mux.Vars(r)
	category := vars["category"]
	id := vars["id"]
	fmt.Printf("Req: %v\n%v\n", category, id)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("View Handler")
	vars := mux.Vars(r)
	title := vars["file"]
	p, err := page.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	page.RenderTemplatePage(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Edit Handler")
	vars := mux.Vars(r)
	title := vars["file"]
	p, err := page.LoadPage(title)
	if err != nil {
		p = &page.Page{Title: title}
	}
	page.RenderTemplatePage(w, "edit", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Save Handler")
	vars := mux.Vars(r)
	title := vars["file"]
	body := r.FormValue("body")
	p := &page.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(customer.GetCustomers()))
}
func customerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//custID, _ := strconv.Atoi(id)
	cust := customer.GetCustomer(id)
	
	w.Write(cust)
}
/*
func customerUpdateHandler(w http.ResponseWriter, r *http.Request) {
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
func customerDeleteHandler(w http.ResponseWriter, r *http.Request) {
	custID, _ := strconv.Atoi(params.ByName("id"))
	cust := customer.DeleteCustomer(custID)
	setHeaders(&w)
	io.WriteString(w, cust)
}
*/
func handleAuthRoutes (w http.ResponseWriter, r *http.Request) {
	log.Println("In auth handler")
	p := &page.Page{Title: "TOKEN", Body: []byte("Token x-Auth will be set in localStorage!!!")}
	page.RenderTemplatePage(w, "getToken", p)
}
//AuthMiddleware checks auth token
type AuthMiddleware struct {
    
}

// NewAuthMiddleware is a struct that has a ServeHTTP method
func NewAuthMiddleware() *AuthMiddleware {
    return &AuthMiddleware{}
}

// The middleware handler
func (l *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
   if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", `POST, GET, OPTIONS,
        	PUT, DELETE`)
		w.Header().Set("Access-Control-Allow-Headers",
			`Accept, Content-Type, Content-Length, Accept-Encoding,
            X-CSRF-Token, Authorization`)
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
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
    // and takes precedence over RemoteAddr Header.Get is case-insensitive
    forward := r.Header.Get("X-Forwarded-For")
    log.Printf("Client's IP: %s\n", ip)
    log.Printf("Client's Port: %s\n", port)
    log.Printf("Client's Forwarded for: %s\n", forward)
	authHeader := r.Header.Get("x-Auth")
	if len(authHeader) != 0 {
  		next(w, r)//call next function on the stack of middleware
	}else {
		http.Redirect(w, r, "/token/getToken", http.StatusTemporaryRedirect)
	}
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/articles/{category}").Subrouter()
	r.HandleFunc("/articles/{category}/order/{sort:(?:asc|desc|new)}", ArticlesCategoryHandler)
	s.HandleFunc("/{id:[0-9]+}", ArticleHandler).Methods("GET").Queries("who", "me")


	//mapping request paths to handlers
	r.HandleFunc("/token/getToken", handleAuthRoutes)
	
	r.HandleFunc("/view/{file:[a-z,A-Z,0-9]+}", viewHandler)
	r.HandleFunc("/edit/{file:[a-z,A-Z,0-9]+}", editHandler)
	saveRouter := mux.NewRouter().PathPrefix("/save").Subrouter()//goes through negroni middleware
	saveRouter.HandleFunc("/{file:[a-z,A-Z,0-9]+}", saveHandler).Methods("POST")
	r.PathPrefix("/save").Handler(negroni.New(
		NewAuthMiddleware(),
		negroni.Wrap(saveRouter),
	))
	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/customers", customersHandler).Methods("GET")
	apiRouter.HandleFunc("/customer/{id:C[0-9]+}", customerHandler).Methods("GET")
	r.PathPrefix("/api").Handler(negroni.New(
		NewAuthMiddleware(),
		negroni.Wrap(apiRouter),
	))
	/*
	
	router.PUT("/api/customer/:id", customerUpdateHandler)
	router.DELETE("/api/customer/:id", customerDeleteHandler)
	*/
	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8070",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
