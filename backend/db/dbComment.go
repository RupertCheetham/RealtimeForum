package db

import (
	"log"
)

// adds a post to the database
func AddCommentToDatabase(username string, parentPostId int, body string) error {
	var likes = 0
	var dislikes = 0
	var whoLiked = ""
	var whoDisliked = ""
	_, err := Database.Exec("INSERT INTO COMMENTS (PostID, Username, Body, Likes, Dislikes, WhoLiked, WhoDisliked) VALUES (?, ?, ?, ?, ?, ?, ?)", parentPostId, username, body, likes, dislikes, whoLiked, whoDisliked)
	if err != nil {
		log.Println("Error adding post to database in AddCommentToDatabase:", err)
	}
	return err
}

// retrieves all posts from database and returns them
func GetCommentsFromDatabase() ([]CommentEntry, error) {

	rows, err := Database.Query("SELECT * FROM COMMENTS ORDER BY Id ASC")
	if err != nil {
		log.Println("Error querying comments from database:", err)
		return nil, err
	}
	defer rows.Close()

	var comments []CommentEntry
	for rows.Next() {

		var comment CommentEntry
		err := rows.Scan(&comment.Id, &comment.ParentPostID, &comment.Username, &comment.Body, &comment.CreationDate, &comment.Likes, &comment.Dislikes, &comment.WhoLiked, &comment.WhoDisliked)
		if err != nil {
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
