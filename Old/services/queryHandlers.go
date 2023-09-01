package forum

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// gets userID from DB
func GetUserIDfrom(table, column, value string) (string, error) {
	var Username string
	query := fmt.Sprintf("SELECT userid FROM %s WHERE %s = ?", table, column)
	row := Database.QueryRow(query, value)
	if err := row.Scan(&Username); err != nil {
		return "", err
	}
	return Username, nil
}

// gets username from DB
func getUsernameFromUserID(inputID string) (string, error) {

	var Username string
	query := fmt.Sprintf("SELECT Username FROM Users WHERE UserID = ?")
	row := Database.QueryRow(query, inputID)
	if err := row.Scan(&Username); err != nil {
		return "", err
	}
	return Username, nil
}

// gets user has from DB
func getUserHash(value string) (string, error) {
	var userHash string
	query := "SELECT password FROM users WHERE username = ?"
	row := Database.QueryRow(query, value)
	if err := row.Scan(&userHash); err != nil {
		return "", err
	}
	return userHash, nil
}

// returns true if the user exists and an error
func UserExists(username, email string) (bool, error) {
	var existingUser string

	if username == "" {
		query := "SELECT email FROM users WHERE email = ?"
		row := Database.QueryRow(query, email)
		err := row.Scan(&existingUser)
		if err == sql.ErrNoRows {
			return false, nil
		} else if err != nil {
			return false, err
		}
	} else if email == "" {
		query := "SELECT username FROM users WHERE username = ?"
		row := Database.QueryRow(query, username)
		err := row.Scan(&existingUser)
		if err == sql.ErrNoRows {
			return false, nil
		} else if err != nil {
			return false, err
		}
	}
	return true, nil

}

// Creates an array of Post(s), to send to the html
func getPostsFromDB() []Post {
	var categoriesString string
	var posts []Post
	rows, err := Database.Query("SELECT * FROM Posts ORDER BY PostID DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		var post Post
		err := rows.Scan(&post.PostID, &post.Username, &post.Img, &post.Body, &categoriesString, &post.CreationDate, &post.Likes, &post.Dislikes, &post.whoLiked, &post.whoDisliked)
		if err != nil {
			log.Fatal(err)
		}

		post.Categories = strings.Split(categoriesString, ",")
		post.CreationDate = formatCreationDate(post.CreationDate)
		stringPostID := strconv.Itoa(post.PostID)
		post.Username, err = getUsernameFromUserID(post.Username)
		post.Comments = GetCommentsFromDB(stringPostID)
		post.CommentCount = len(post.Comments)

		if err != nil {

			log.Println(err)
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return posts
}

// Creates an array of Comment(s), to send to the html
func GetCommentsFromDB(postID string) []Comment {

	var comments []Comment
	rows, err := Database.Query("SELECT * FROM Comments WHERE PostID = ?", postID)
	if err != nil {
		log.Fatal(err, "in GetCommentsFromDB")
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.Username, &comment.Body, &comment.CreationDate, &comment.Likes, &comment.Dislikes, &comment.whoLiked, &comment.whoDisliked)
		if err != nil {
			log.Fatal(err)
		}

		comment.CreationDate = formatCreationDate(comment.CreationDate)
		comment.Username, err = getUsernameFromUserID(comment.Username)
		if err != nil {
			log.Println(err, "in GetCommentsFromDB")
		}

		comments = append(comments, comment)

	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return comments
}

// formats SQL date to one that is more visually appealing (for getPostsFromDB)
func formatCreationDate(CreationDate string) string {

	convertedDate, err := time.Parse(time.RFC3339, CreationDate)
	if err != nil {
		// Handle parsing error
		log.Println("Failed to parse CreationDate:", err)
		return "Failed to parse CreationDate"
	}
	// The database is an hour behind, this tweaks that
	FurtherConvertedDate := convertedDate.Local()
	// Formats the time to a more pleasing version
	CreationDate = FurtherConvertedDate.Format("15:04 02-01-2006")
	return CreationDate

}
