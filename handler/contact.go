package handler

import (
	"html/template"
	"net/http"

	"github.com/gorilla/schema"
)

//basically will handle the functionalities of the form in contacts page
//from the link with contact.html

func Contact(tmpl *template.Template) http.HandlerFunc {
	type formData struct {
		email   string `schema: "email"`
		subject string `schema: "subject"`
		message string `schema: "message"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl.ExecuteTemplate(w, "contact.html", nil)
		}
		if r.Method == http.MethodPost {
			r.ParseForm()
			var input formData
			var decoder = schema.NewDecoder()
			if err := decoder.Decode(&input, r.PostForm); err != nil {
				http.Error(w, "Bad form", http.StatusBadRequest)
				return
			}
		}
	}

}
