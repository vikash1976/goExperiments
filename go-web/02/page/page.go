package page

import (
    "io/ioutil"
    "net/http"
    "html/template"
    
)
//Page Type with Title and Body
type Page struct {
    Title string
    Body []byte
}
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))


//Save function saves the body content as file
func (p *Page) Save()  error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}
//LoadPage loads the provide file content as page
func LoadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err:= ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
    
}

//RenderTemplatePage function renders the template given as 2nd arg
func RenderTemplatePage(w http.ResponseWriter, tmpl string, p *Page) {
   
    err := templates.ExecuteTemplate(w, tmpl + ".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    } 
}
