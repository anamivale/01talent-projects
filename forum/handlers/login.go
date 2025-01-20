package handlers

import (
	"html/template"
	"net/http"

	"github.com/anamivale/01talent-projects.git/middlewares"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == http.MethodPost {
		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("password")

		MidlwarEerr := middlewares.Login(email, password)
		if MidlwarEerr != nil {
			t, err := template.ParseFiles("./templates/login.html")
			if err != nil {
				http.Error(w, "could not load template", http.StatusInternalServerError)
				return
			}

			t.Execute(w, MidlwarEerr.Error())
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	t, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		http.Error(w, "could not load template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
