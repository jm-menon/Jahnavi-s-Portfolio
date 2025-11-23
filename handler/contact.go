package handler

import (
	"html/template"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jm-menon/Jahnavi-s-Portfolio/internal/mail"
)

//basically will handle the functionalities of the form in contacts page
//from the link with contact.html

func Contact(tmpl *template.Template) http.HandlerFunc {
	type formData struct {
		Email   string `schema: "email"`
		Subject string `schema: "subject"`
		Message string `schema: "message"`
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
			if err := mail.SendContact(input.Email, input.Subject, input.Message); err != nil {
				tmpl.ExecuteTemplate(w, "contact.html", map[string]string{
					"Error": "Failed to send. Try again.",
				})
				return
			}
			tmpl.ExecuteTemplate(w, "contact.html", map[string]string{
				"Success": "Emali sent successfully! I will reply soon!!"})
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
