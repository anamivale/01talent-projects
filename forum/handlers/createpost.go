package handlers

import (
	"net/http"
	"text/template"

	"github.com/anamivale/01talent-projects.git/middlewares"
)

func Createpost(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == http.MethodPost {
		r.ParseForm()
		content := r.FormValue("post")
		enta := r.FormValue("Entertainment")
		ed := r.FormValue("Education")
		fashion := r.FormValue("Fashion")
		sports := r.FormValue("Sports")

		category := ""

		if enta != "" {
			category = enta
		}
		if ed != "" {
			category += " " + ed
		}
		if fashion != "" {
			category += " " + fashion
		}
		if sports != "" {
			category += " " + sports
		}
		id := middlewares.GenerateId()
		middlewares.CreatePost(id,content, "valeria", 0, 0, category)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	t, err := template.ParseFiles("./templates/post.html")
	if err != nil {
		http.Error(w, "could not load template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
