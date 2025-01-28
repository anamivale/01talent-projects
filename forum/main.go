package main

import (
	"fmt"
	"net/http"

	"github.com/anamivale/01talent-projects.git/handlers"
	"github.com/anamivale/01talent-projects.git/middlewares"
)

func main() {
	middlewares.DbConn()

	middlewares.CreateUsersTable()
	middlewares.CreatePostsTable()
	middlewares.CreateCommentTable()
	middlewares.Like()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/post", handlers.Createpost)
	http.HandleFunc("/like", handlers.LikeHandler)
	http.HandleFunc("/logout", handlers.Logout)

	fmt.Println("Server running at http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err.Error())
	}

	middlewares.GenerateId()
}
