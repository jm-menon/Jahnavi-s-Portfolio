package handler

import (
	"html/template"
	"net/http"
)

func HTML(page string, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, page, nil)
	}
}
