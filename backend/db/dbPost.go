package db

import (
	"log"
	"realtimeForum/utils"
	"strings"
)

// adds a post to the database
func AddPostToDatabase(userID int, img string, body string, categories string) error {

	_, err := Database.Exec("INSERT INTO POSTS (UserId, Img, Body, Categories) VALUES (?, ?, ?, ?)", userID, img, body, categories)
	if err != nil {
		utils.HandleError("Error adding post to database in addPostToDatabase:", err)
		log.Println("Error adding post to database in addPostToDatabase:", err)
	}
	return err
}

// retrieves all posts from database and returns them
func GetPostFromDatabase() ([]PostEntry, error) {
	query := `
	SELECT p.Id, u.Username, p.Img, p.Body, p.Categories, p.CreationDate, p.ReactionID,
	COALESCE(pr.Likes, 0) AS Likes, COALESCE(pr.Dislikes, 0) AS Dislikes
FROM POSTS p
LEFT JOIN POSTREACTIONS pr ON p.ReactionID = pr.Id
LEFT JOIN USERS u ON p.UserId = u.Id
ORDER BY p.Id DESC;

`

	rows, err := Database.Query(query)
	if err != nil {
		utils.HandleError("Error querying posts with likes and dislikes from database:", err)
		log.Println("Error querying posts with likes and dislikes from database:", err)
		return nil, err
	}
	defer rows.Close()

	var posts []PostEntry
	for rows.Next() {
		var post PostEntry
		var categoriesString string
		err := rows.Scan(&post.Id, &post.Username, &post.Img, &post.Body, &categoriesString, &post.CreationDate, &post.ReactionID, &post.Likes, &post.Dislikes)
		if err != nil {
			utils.HandleError("Error scanning row from database:", err)
			return nil, err
		}
		post.Categories = strings.Split(categoriesString, ",")
		posts = append(posts, post)
	}

	return posts, nil
}
