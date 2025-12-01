package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jm-menon/Jahnavi-s-Portfolio/handler"
)

var decoder = schema.NewDecoder()

func NewServer() http.Handler {

	mux := http.NewServeMux()

	//mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("server")))) //ask an explanation

	tmpl := template.Must(template.ParseGlob("pages/*.html")) //ask for a detailed explanation

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			log.Println("404 Error", r.URL.Path)
			return
		}
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})
	mux.HandleFunc("/about", handler.HTML("about.html", tmpl))
	mux.HandleFunc("/blogs", handler.HTML("blogs.html", tmpl))
	mux.HandleFunc("/projects", handler.HTML("projects.html", tmpl))
	mux.HandleFunc("/resume", handler.PDF("assets/Software_Engineer_Jahnavi_Menon.pdf"))
	mux.HandleFunc("/contact", handler.Contact(tmpl))

	return loggingMiddleware(recoveryMiddleware(mux))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
