package server

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("html/index.html", "html/about.html"))

func render(w http.ResponseWriter, page string) {
	err := templates.ExecuteTemplate(w, page+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/" {
		render(w, "index")
		return
	}
	http.NotFound(w, r)
	return
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "about")
}

func init() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)
}
