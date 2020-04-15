package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("index.html"))

type Page struct {}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", &Page{})
}

func featurePolicy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Feature-Policy", "sync-xhr 'none'")
	renderTemplate(w, "index", &Page{})
}

func xhr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("synchronous xhr request is executed."))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/feature-policy", featurePolicy)
	http.HandleFunc("/xhr", xhr)
	log.Fatal(http.ListenAndServe(":8080", nil))
}