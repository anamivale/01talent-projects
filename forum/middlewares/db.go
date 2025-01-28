package middlewares

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var err error

var db *sql.DB

func DbConn() {
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func CreateUsersTable() {
	table := `
	CREATE TABLE IF NOT EXISTs users(
	username TEXT NOT NULL UNIQUE PRIMARY KEY,
	email TEXT  UNIQUE,
	password TEXT NOT NULL
)
`
	_, err := db.Exec(table)
	CheckErr(err, "table not created")
}

func Signup(username, email, password string) {
	querystring := `
	INSERT INTO users (username,email, password) VALUES (?,?,?)`

	_, err = db.Exec(querystring, username, email, password)
	if err != nil {
		CheckErr(err, "could not insert")
	}
}

func Login(email, password string) error {
	var hashedPass string
	queryStmt := `SELECT password FROM users WHERE email = ?`
	err = db.QueryRow(queryStmt, email).Scan(&hashedPass)

	if err == sql.ErrNoRows {
		return errors.New("user not found")
	}
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("incorect password")
	}

	return nil
}

func CreatePostsTable() {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS posts(
		id TEXT PRIMARY KEY NOT NULL UNIQUE ,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		like INTEGER,
		dislike INTEGER,
		category TEXT,
		name  TEXT NOT NULL,
		FOREIGN KEY (name) REFERENCES users(username) 

	)
	`

	_, err = db.Exec(sqlQuery)
	if err != nil {
		fmt.Println(err.Error(), 3)
		return
	}
}

func CreatePost(id, content, name string, like, dislike int, category string) {
	queryStmt := `
	INSERT INTO posts(id,content, like, dislike,category, name) VALUES (?,?,?,?,?,?)
	`
	_, err = db.Exec(queryStmt, id, content, like, dislike, category, name)
	if err != nil {
		fmt.Println(err.Error(), 2)
		return
	}
}

func GetPosts() []*Post {
	var posts []*Post

	queryStmt := `
	SELECT id, name, content, like, dislike, created_at FROM posts`

	rows, err := db.Query(queryStmt)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		// Create a new instance of Post for each row
		post := &Post{}
		err := rows.Scan(&post.Id, &post.Name, &post.Content, &post.Like, &post.Dislike, &post.Created_at)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		post.PostComments = GetComments(post.Id)
		// Append the new instance to the slice
		posts = append(posts, post)
	}

	return posts
}

func CreateCommentTable() {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS comments(
		id 	TEXT PRIMARY KEY  NOT NULL,
		name TEXT NOT NULL,
		content TEXT NOT NULL,
		like INTEGER Default 0,
		dislike INTEGER Default 0,
		post_id TEXT NOT NULL,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
		FOREIGN KEY (name) REFERENCES users(username) ON DELETE CASCADE
	)
	`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		fmt.Println(err.Error(), "1")
		return
	}
}

func CreateComment(id, name, content, post_id string) {
	stmt := `
	INSERT INTO comments(id,name, content, post_id) VALUES (?,?,?,?)
	`
	_, err = db.Exec(stmt, id, name, content, post_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func GetComments(post_id string) []Comments {
	var comments []Comments
	qstmt := `
	SELECT name, content, like , dislike FROM comments WHERE post_id = ?
	`

	rows, err := db.Query(qstmt, post_id)
	if err != nil {
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		comment := Comments{}

		err = rows.Scan(&comment.Name, &comment.Content, &comment.Like, &comment.Dislike)
		if err != nil {
			fmt.Println("could not unpack")
			return nil
		}
		comments = append(comments, comment)
	}

	return comments
}

func Like() {
	qStmt := `
	CREATE TABLE IF NOT EXISTS likes (
		post_id TEXT NOT NULL,  
		user_id TEXT NOT NULL
	);
	`
	_, err := db.Exec(qStmt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
