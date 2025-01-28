package middlewares

import (
	"fmt"
	"log"
)

func AddLike(postID, userID string) error {
	query := `INSERT INTO likes (post_id, user_id) VALUES (?, ?)`
	_, err := db.Exec(query, postID, userID)

	return err
}

func RemoveLike(postID, userID string) error {
	query := `DELETE FROM likes WHERE post_id = $1 AND user_id = $2`
	_, err := db.Exec(query, postID, userID)
	return err
}

func GetLikeCount(postID string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM likes WHERE post_id = $1`
	err := db.QueryRow(query, postID).Scan(&count)
	return count, err
}

func CheckIfAlreadyLiked(postid, userid string) bool {
	qStmt := `
    SELECT EXISTS (
        SELECT *
        FROM likes 
        WHERE post_id = ? AND user_id = ?
    )`

	var exists bool
	err := db.QueryRow(qStmt, postid, userid).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if already liked: %v", err)
		return false
	}

	return exists
}

func UpdateLikeCount(Likes int, postId string) {
	qStmt := `
	UPDATE posts SET like = ? WHERE id = ? 
	`
	_, err := db.Exec(qStmt, Likes, postId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
