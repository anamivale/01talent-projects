package handlers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Could not load index template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
