package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/vikash1976/goExperiments/go-web/02/page"
	"net/http"
	"regexp"
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

func main() {
	router := httprouter.New()
	router.GET("/view/:file", makeHandler(viewHandler))
	router.GET("/edit/:file", makeHandler(editHandler))
	router.POST("/save/:file", makeHandler(saveHandler))
	/*http.HandleFunc("/view/",  makeHandler(viewHandler))
	  http.HandleFunc("/edit/",  makeHandler(editHandler))
	  http.HandleFunc("/save/",  makeHandler(saveHandler))*/
	http.ListenAndServe(":8070", router)
}
