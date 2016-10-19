package page

import (
	
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

//Page Type with Title and Body
type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
//var validPath = regexp.MustCompile("([a-zA-Z0-9]+)$")

//Save function saves the body content as file
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//LoadPage loads the provide file content as page
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	fmt.Printf("Reading file: %s\n", filename)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

}

//GetTitle gets the title after validation per validPath
/*func GetTitle(w http.ResponseWriter, r *http.Request, params httprouter.Params) (string, error) {
	m := validPath.FindStringSubmatch(params.ByName("file"))
	if m == nil {
		//http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[1], nil // The title is the second subexpression.
}*/

//RenderTemplatePage function renders the template given as 2nd arg
func RenderTemplatePage(w http.ResponseWriter, tmpl string, p *Page) {

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
