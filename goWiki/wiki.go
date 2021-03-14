package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)


// template caching
var templates = template.Must(template.ParseFiles("view.html", "edit.html"))

type Page struct {
	Title string
	Body  []byte
}


func (page *Page) save() error {
	var fileName = page.Title + ".txt"
	return ioutil.WriteFile(fileName, page.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}


func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", page)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(w, "edit", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")

	page := &Page{title, []byte(body) }
	err := page.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/view/" + title, http.StatusFound)

}

func renderTemplate(w http.ResponseWriter, tmpl string, page *Page) {

	err := templates.ExecuteTemplate(w, tmpl + ".html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}