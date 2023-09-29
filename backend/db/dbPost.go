package db

import (
	"encoding/json"
	"log"
	"realtimeForum/utils"
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
	rows, err := Database.Query("SELECT * FROM POSTS ORDER BY Id DESC")
	if err != nil {
		utils.HandleError("Error querying POSTS from database:", err)
		log.Println("Error querying POSTS from database:", err)
		return nil, err
	}
	defer rows.Close()

	var posts []PostEntry
	for rows.Next() {
		var post PostEntry
		err := rows.Scan(&post.Id, &post.UserId, &post.Img, &post.Body, &post.Categories, &post.CreationDate, &post.ReactionID)
		if err != nil {
			utils.HandleError("Error scanning row from database:", err)
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func ConvertPostsToJSON() ([]byte, error) {

	posts, _ := GetPostFromDatabase()
	// Marshal the array of PostEntry structs into JSON
	jsonPosts, err := json.Marshal(posts)
	if err != nil {
		return nil, err
	}
	return jsonPosts, nil
}
