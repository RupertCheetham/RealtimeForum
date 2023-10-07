package db

import (
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
	query := `
	SELECT p.Id, p.UserId, p.Img, p.Body, p.Categories, p.CreationDate, p.ReactionID,
		   COALESCE(pr.Likes, 0) AS Likes, COALESCE(pr.Dislikes, 0) AS Dislikes
	FROM POSTS p
	LEFT JOIN POSTREACTIONS pr ON p.ReactionID = pr.Id
	ORDER BY p.Id DESC
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
		err := rows.Scan(&post.Id, &post.UserId, &post.Img, &post.Body, &post.Categories, &post.CreationDate, &post.ReactionID, &post.Likes, &post.Dislikes)
		if err != nil {
			utils.HandleError("Error scanning row from database:", err)
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// retrieves newest post from database and returns it
func GetNewestPost() (PostEntry, error) {
	query := `
        SELECT
            Id,
            UserID,
            Img,
            Body,
            Categories,
            CreationDate,
            ReactionID,
            0 AS postLikes,
            0 AS postDislikes
        FROM POSTS
        ORDER BY Id DESC
        LIMIT 1;
    `
	var newestPost PostEntry

	err := Database.QueryRow(query).Scan(
		&newestPost.Id,
		&newestPost.UserId,
		&newestPost.Img,
		&newestPost.Body,
		&newestPost.Categories,
		&newestPost.CreationDate,
		&newestPost.ReactionID,
		&newestPost.Likes,
		&newestPost.Dislikes,
	)

	if err != nil {
		return newestPost, err
	}

	return newestPost, nil
}
