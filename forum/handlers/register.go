package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/anamivale/01talent-projects.git/middlewares"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		userName := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmpass := r.FormValue("confirmpass")
		fmt.Println(userName, email, password, confirmpass)
		if confirmpass != password {
			t, err := template.ParseFiles("./templates/register.html")
			if err != nil {
				http.Error(w, "Could not load register template", http.StatusInternalServerError)
				return
			}
			data := "Passwords do not match"
			t.Execute(w, data)
			return
		}
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		middlewares.CheckErr(err, "Could not encrypt the password")

		middlewares.Signup(userName, email, string(hashedPass))

		fmt.Println("User Registered")

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		// Serve the registration form if not a POST request
		t, err := template.ParseFiles("./templates/register.html")
		if err != nil {
			http.Error(w, "Could not load register template", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)

	}
}
