package forum

import (
	"log"
	"strconv"
	"strings"
)

// Finds and returns Posts made by the User
func FilterPostsByUserID(UserID string) []Post {

	var filteredPosts []Post

	// Prepare a statement to get all posts by the current user
	rows, err := Database.Query("SELECT * FROM Posts WHERE UserID = ?", UserID)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the results and print each post
	for rows.Next() {
		var post Post
		var categoriesString string
		err := rows.Scan(&post.PostID, &post.Username, &post.Img, &post.Body, &categoriesString, &post.CreationDate, &post.Likes, &post.Dislikes, &post.whoLiked, &post.whoDisliked)
		if err != nil {
			log.Fatal(err)
		}
		post.Categories = strings.Split(categoriesString, ",")
		post.CreationDate = formatCreationDate(post.CreationDate)
		stringPostID := strconv.Itoa(post.PostID)
		post.Username, _ = getUsernameFromUserID(post.Username)
		post.Comments = GetCommentsFromDB(stringPostID)
		post.CommentCount = len(post.Comments)
		filteredPosts = append(filteredPosts, post)
	}
	return filteredPosts
}

// Finds the Users Liked Posts and returns them as an array of Post
func FilterLikedPostsByUserID(UserID string) []Post {

	var filteredLikedPosts []Post

	// Prepare a statement to get all posts by the current user
	rows, err := Database.Query("SELECT * FROM Posts ORDER BY PostID DESC")
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the results and print each post
	for rows.Next() {
		var post Post
		var categoriesString string
		err := rows.Scan(&post.PostID, &post.Username, &post.Img, &post.Body, &categoriesString, &post.CreationDate, &post.Likes, &post.Dislikes, &post.whoLiked, &post.whoDisliked)
		if err != nil {
			log.Fatal(err)
		}

		splitLiked := strings.Split(post.whoLiked, ",")
		for _, idAccount := range splitLiked {
			if idAccount == UserID {
				post.Categories = strings.Split(categoriesString, ",")
				post.CreationDate = formatCreationDate(post.CreationDate)
				stringPostID := strconv.Itoa(post.PostID)
				post.Username, err = getUsernameFromUserID(post.Username)
				post.Comments = GetCommentsFromDB(stringPostID)
				post.CommentCount = len(post.Comments)
				filteredLikedPosts = append(filteredLikedPosts, post)
			}
		}

	}

	return filteredLikedPosts
}

// Finds and returns Posts made by the User
func FilterPostsByCategory(categoryRequest string) []Post {

	var filteredPosts []Post

	// Prepare a statement to get all posts by the current user
	rows, err := Database.Query("SELECT * FROM Posts")
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the results and print each post
	for rows.Next() {
		var categoriesString string
		var post Post
		err := rows.Scan(&post.PostID, &post.Username, &post.Img, &post.Body, &categoriesString, &post.CreationDate, &post.Likes, &post.Dislikes, &post.whoLiked, &post.whoDisliked)
		if err != nil {
			log.Fatal(err)
		}

		categories := strings.Split(categoriesString, ",")
		for _, category := range categories {

			if category == categoryRequest {
				post.Categories = categories
				post.CreationDate = formatCreationDate(post.CreationDate)
				stringPostID := strconv.Itoa(post.PostID)
				post.Username, _ = getUsernameFromUserID(post.Username)
				post.Comments = GetCommentsFromDB(stringPostID)
				post.CommentCount = len(post.Comments)
				filteredPosts = append(filteredPosts, post)
				continue // Skip processing this row and move to the next iteration
			}
		}
	}
	return filteredPosts
}
