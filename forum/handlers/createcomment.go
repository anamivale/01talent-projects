package handlers

import (
	"fmt"
	"net/http"

	"github.com/anamivale/01talent-projects.git/middlewares"
)


func CreateComment(w http.ResponseWriter, r *http.Request)  {
		id :=  middlewares.GenerateId()
		cookie,err := middlewares.GetCookies(r,"session-cookies")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		name := cookie.Value

		content := r.FormValue("content")
		postid := r.FormValue("postid")

		middlewares.CreateComment(id, name, content,postid)

		
	

		
}