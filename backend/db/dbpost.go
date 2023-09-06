package db

import "log"

// adds a post to the database
func AddPostToDatabase(nickname string, img string, body string, categories string) error {
	var likes = 0
	var dislikes = 0
	var whoLiked = ""
	var whoDisliked = ""
	_, err := Database.Exec("INSERT INTO Posts (Nickname, Img, Body, Categories, Likes, Dislikes, WhoLiked, WhoDisliked) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", nickname, img, body, categories, likes, dislikes, whoLiked, whoDisliked)
	if err != nil {
		log.Println("Error adding post to database in addPostToDatabase:", err)
	}
	return err
}

// retrieves all posts from database and returns them
func GetPostFromDatabase() ([]PostEntry, error) {
	rows, err := Database.Query("SELECT Id, Nickname, Img, Body, Categories, CreationDate, Likes, Dislikes, WhoLiked, WhoDisliked FROM Posts ORDER BY Id ASC")
	if err != nil {
		log.Println("Error querying posts from database:", err)
		return nil, err
	}
	defer rows.Close()

	var posts []PostEntry
	for rows.Next() {
		var post PostEntry
		err := rows.Scan(&post.Id, &post.Nickname, &post.Img, &post.Body, &post.Categories, &post.CreationDate, &post.Likes, &post.Dislikes, &post.WhoLiked, &post.WhoDisliked)
		if err != nil {
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
