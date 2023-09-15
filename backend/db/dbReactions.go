package db

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// if Reaction > 0 {UPDATE} else {INSERT}

func AddReactionToDatabase(tableName string, userID int, reaction string) {

	likes := 0
	dislikes := 0
	whoLiked := ""
	whoDisliked := ""
	query := ""

	if reaction == "like" {
		likes = 1
		whoLiked = whoLiked + strconv.Itoa(userID)
		query = fmt.Sprintf("INSERT INTO %s (Likes, WhoLiked) VALUES (?, ?)", tableName)
	} else {
		dislikes = 1
		whoDisliked = whoDisliked + strconv.Itoa(userID)
		query = fmt.Sprintf("INSERT INTO %s (Dislikes, WhoDisliked) VALUES (?, ?)", tableName)
	}

	_, err := Database.Exec(query, likes, dislikes, whoLiked, whoDisliked)
	if err != nil {
		log.Println("Error adding comment to database in AddReactionToDatabase:", err)
	}
}

func UpdateReactionInDatabase(tableName string, rowID int, userID int, reaction string) {

	var likes int
	var dislikes int
	var whoLiked string
	var whoDisliked string

	// Construct the SELECT query to retrieve current values
	selectQuery := fmt.Sprintf("SELECT Likes, Dislikes, WhoLiked, WhoDisliked FROM %s WHERE Id = ?", tableName)

	err := Database.QueryRow(selectQuery, rowID).Scan(&likes, &dislikes, &whoLiked, &whoDisliked)
	if err != nil {
		log.Println("Error retrieving values from the database:", err)
		return
	}

	// Modify the values based on the reaction
	if reaction == "like" {
		likes, dislikes, whoLiked, whoDisliked = ReactionAdjuster(userID, likes, dislikes, whoLiked, whoDisliked)

	} else {
		dislikes, likes, whoDisliked, whoLiked = ReactionAdjuster(userID, dislikes, likes, whoDisliked, whoLiked)
	}

	// Construct the UPDATE query to update the table
	updateQuery := fmt.Sprintf("UPDATE %s SET Likes = ?, Dislikes = ?, WhoLiked = ?, WhoDisliked = ? WHERE Id = ?", tableName)

	_, err = Database.Exec(updateQuery, likes, dislikes, whoLiked, whoDisliked, rowID)
	if err != nil {
		log.Println("Error updating values in the database:", err)
		return
	}
}

// Tweaks the values of likes/dislikes as needed
func ReactionAdjuster(userID int, value1 int, value2 int, who1 string, who2 string) (int, int, string, string) {

	userIDstring := strconv.Itoa(userID)
	// checks to see if the like (or dislike) is a repeat action, if it is then returns values unchanged
	splitA := strings.Split(who1, ",")
	for _, idAccount := range splitA {
		if idAccount == userIDstring {
			return value1, value2, who1, who2
		}
	}

	// checks to see if the oppsite action has already taken place, if it has then returns removes the action from the db
	splitB := strings.Split(who2, ",")
	for i, idAccount := range splitB {
		if idAccount == userIDstring {
			value2 = value2 - 1
			// removes userID so they're no longer on the list as having performed the opposite action
			splitB = append(splitB[:i], splitB[i+1:]...)
		}
	}

	// performs the action (like/dislike)
	who2 = strings.Join(splitB, ",")
	who2 = strings.TrimPrefix(who2, ",")

	// adds userID to whoLiked or whoDisliked
	who1 = who1 + "," + userIDstring
	who1 = strings.TrimPrefix(who1, ",")
	value1 = value1 + 1

	return value1, value2, who1, who2

}
