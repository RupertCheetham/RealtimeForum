package db

import (
	"log"
	"realtimeForum/utils"
)

// adds a post to the database
func AddCommentToDatabase(parentPostID int, userID int, body string) error {
	_, err := Database.Exec("INSERT INTO COMMENTS (PostID, UserID, Body) VALUES (?, ?, ?)", parentPostID, userID, body)
	if err != nil {
		utils.HandleError("Error adding comment to database in AddCommentToDatabase:", err)
		log.Println("Error adding comment to database in AddCommentToDatabase:", err)
	}
	return err
}

// retrieves all posts from database and returns them
func GetCommentsFromDatabase() ([]CommentEntry, error) {

	rows, err := Database.Query("SELECT * FROM COMMENTS ORDER BY Id ASC")
	if err != nil {
		utils.HandleError("Error querying comments from database:", err)
		log.Println("Error querying comments from database:", err)
		return nil, err
	}
	defer rows.Close()

	var comments []CommentEntry
	for rows.Next() {

		var comment CommentEntry
		err := rows.Scan(&comment.Id, &comment.ParentPostID, &comment.UserId, &comment.Body, &comment.CreationDate, &comment.ReactionID)
		if err != nil {
			utils.HandleError("Error scanning row from database:", err)
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
