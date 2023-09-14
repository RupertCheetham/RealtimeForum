package db

import (
	"log"
)

// adds a post to the database
func AddCommentToDatabase(userID int, parentPostID int, body string) error {
	var reaction = 0
	_, err := Database.Exec("INSERT INTO COMMENTS (PostID, UserId, Body, Reaction) VALUES (?, ?, ?, ?)", parentPostID, userID, body, reaction)
	if err != nil {
		log.Println("Error adding comment to database in AddCommentToDatabase:", err)
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
		err := rows.Scan(&comment.Id, &comment.ParentPostID, &comment.UserId, &comment.Body, &comment.CreationDate, &comment.Reaction)
		if err != nil {
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
