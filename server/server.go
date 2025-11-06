package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func NewServer() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("server")))) //ask an explanation

	tmpl := template.Must(template.ParseGlob("/pages/*.html")) //ask for a detailed explanation

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			log.Println("404 Error", r.URL.Path)
			return
		}
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	genericHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Another Page")
		w.Write([]byte("Another Page"))
	}
	mux.HandleFunc("/about", genericHandler)
	mux.HandleFunc("/projects", genericHandler)
	mux.HandleFunc("/contact", genericHandler)
	mux.HandleFunc("/blog", genericHandler)

	return mux

}
