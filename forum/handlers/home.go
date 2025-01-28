package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/anamivale/01talent-projects.git/middlewares"
)

func Home(w http.ResponseWriter, r *http.Request) {
	isLoggedIn := false
	cookie, errC := middlewares.GetCookies(r, "session-cookies")

	if errC == nil && cookie.Value != "" {
		isLoggedIn = true
	}

	if r.Method == http.MethodPost {
		if errC != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		id := middlewares.GenerateId()
		cookie, err := middlewares.GetCookies(r, "session-cookies")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		name := cookie.Value

		content := r.FormValue("content")
		postid := r.FormValue("postid")

		middlewares.CreateComment(id, name, content, postid)
	}

	post := middlewares.GetPosts()

	// Combine posts and authentication status into a map
	data := map[string]interface{}{
		"Posts":      post,
		"IsLoggedIn": isLoggedIn,
	}

	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Could not load index template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
