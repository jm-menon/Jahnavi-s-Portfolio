package main

import (
	"log"
	"net/http"
)

func main() {

	//fs:= http.FileServer(http.Dir("pages"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			http.ServeFile(w, r, "pages/index.html")
		case "/about":
			http.ServeFile(w, r, "pages/about.html")
		case "/blog":
			http.ServeFile(w, r, "pages/blogs.html")
		case "/projects":
			http.ServeFile(w, r, "pages/projects.html")
		default:
			http.NotFound(w, r)
		}

	})
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
