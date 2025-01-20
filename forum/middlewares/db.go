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

func CreatUser() {
	querystring := `
	INSERT INTO users (username,email, password) VALUES (?,?,?)`

	_, err = db.Exec(querystring, "username", "email@gmail.com", "password")
	if err != nil {
		CheckErr(err, "could not insert")
	}
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
	fmt.Println(hashedPass)
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
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		like INTEGER,
		dislike INTEGER,
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

func CreatePost() {
	queryStmt := `
	INSERT INTO posts(content, like, dislike, name) VALUES (?,?,?,?)
	`
	_, err = db.Exec(queryStmt, "welcome to posts table", 0, 0, "username")
	if err != nil {
		fmt.Println(err.Error(), 2)
		return
	}
}

func CreateCommentTable() {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS comments(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		content TEXT NOT NULL,
		like INTEGER,
		dislike INTEGER,
		post_id INTEGER NOT NULL,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
		FOREIGN KEY (name) REFERENCES users(username) 
	)
	`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		fmt.Println(err.Error(), "1")
		return
	}
}

func CreateComment() {
	stmt := `
	INSERT INTO comments(name, content, like, dislike, post_id) VALUES (?,?,?,?,?)
	`
	_, err = db.Exec(stmt, "username", "welcome to comments section", 0, 0, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
