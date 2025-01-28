package middlewares

import "time"

type Post struct {
	Name       string
	Content    string
	Like       int
	Dislike    int
	Created_at string
	Id         string
	PostComments []Comments
}

type Likes struct {
	ID        int
	PostID    string
	UserID    string
	CreatedAt time.Time
}

type Comments struct {
	Name    string
	Content string
	Like    int
	Dislike int
}
