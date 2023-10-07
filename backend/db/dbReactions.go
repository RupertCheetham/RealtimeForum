package db

import (
	"fmt"
	"log"
	"realtimeForum/utils"
	"strconv"
	"strings"
)

// Adds a reaction to the database, if post/comment doesn't have any.
// Updates the parent post/comment's ReactionID to relevant reactiontable's ID
func AddReactionToDatabase(reactionParentClass string, parentID int, userID int, reaction string) {

	fmt.Println(reactionParentClass, parentID, reaction)
	// Fills in the table details, depending upon if the reaction's parent class is "post" or "comment"
	tableName := ""
	postOrCommentTable := ""
	if reactionParentClass == "post" {
		postOrCommentTable = "POSTS"
		tableName = "POSTREACTIONS"
	} else {
		postOrCommentTable = "COMMENTS"
		tableName = "COMMENTREACTIONS"
	}

	likes := 0
	dislikes := 0
	whoLiked := ""
	whoDisliked := ""

	// if the reaction was "like" then adds a like
	if reaction == "like" {
		likes = 1
		whoLiked = whoLiked + strconv.Itoa(userID)
		// if the reaction was "dislike" then adds a dislike instead
	} else if reaction == "dislike" {
		dislikes = 1
		whoDisliked = whoDisliked + strconv.Itoa(userID)
	} else {
		// Just in case I make a type in the HTML
		log.Println("Invalid reaction in AddReactionToDatabase:", reaction)
		utils.HandleError("Invalid reaction in AddReactionToDatabase:", nil)
	}

	// Insert the reaction into the tableName table
	query := fmt.Sprintf("INSERT INTO %s (Likes, Dislikes, WhoLiked, WhoDisliked) VALUES (?, ?, ?, ?)", tableName)
	_, err := Database.Exec(query, likes, dislikes, whoLiked, whoDisliked)
	if err != nil {
		utils.HandleError("Error adding reaction to database in AddReactionToDatabase:", err)
		log.Println("Error adding reaction to database in AddReactionToDatabase:", err)
	}

	// Update the ReactionID of the specified post/comment (e.g. post 5),
	// but first needs to decide whether to update posts or comments

	updateQuery := fmt.Sprintf("UPDATE %s SET ReactionID = (SELECT Id FROM %s ORDER BY Id DESC LIMIT 1) WHERE Id = ?", postOrCommentTable, tableName)
	_, err = Database.Exec(updateQuery, parentID)
	if err != nil {
		utils.HandleError("Error updating ReactionID in database in AddReactionToDatabase:", err)
		log.Println("Error updating ReactionID in AddReactionToDatabase:", err)
	}
}

// Updates values already in the reaction table
func UpdateReactionInDatabase(reactionParentClass string, rowID int, userID int, reaction string) {

	tableName := ""
	if reactionParentClass == "post" {
		tableName = "POSTREACTIONS"
	} else {
		tableName = "COMMENTREACTIONS"
	}

	var likes int
	var dislikes int
	var whoLiked string
	var whoDisliked string

	// Construct the SELECT query to retrieve current values
	selectQuery := fmt.Sprintf("SELECT Likes, Dislikes, WhoLiked, WhoDisliked FROM %s WHERE Id = ?", tableName)

	err := Database.QueryRow(selectQuery, rowID).Scan(&likes, &dislikes, &whoLiked, &whoDisliked)
	if err != nil {
		utils.HandleError("Error retrieving values from REACTIONS in UpdateReactionInDatabase:", err)
		log.Println("Error retrieving values from REACTIONS in UpdateReactionInDatabase:", err)
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
		utils.HandleError("Error updating values in REACTIONS in UpdateReactionInDatabase:", err)
		log.Println("Error updating values in REACTIONS in UpdateReactionInDatabase:", err)
		return
	}
}

// Tweaks the values of likes/dislikes as needed
func ReactionAdjuster(userID int, actionToIncrement int, oppositeAction int, actionUsers string, oppositeActionUsers string) (int, int, string, string) {

	userIDstring := strconv.Itoa(userID)

	// If the reaction is a repeat of a previous reaction then nothing changes
	if actionAlreadyPerformed(userIDstring, actionToIncrement, oppositeAction, actionUsers, oppositeActionUsers) {
		return actionToIncrement, oppositeAction, actionUsers, oppositeActionUsers
	}

	oppositeActionAlreadyPerformed, user := oppositeActionChecker(userIDstring, oppositeActionUsers)
	// If the User has already performed the opposite reaction then undoes the opposite reaction
	if oppositeActionAlreadyPerformed {
		oppositeAction = oppositeAction - 1
		oppositeActionUsers = removeUserIDFromOppositeActionUsers(oppositeActionUsers, user)
	}

	// Finally, performs the action (like/dislike)
	actionUsers = addUsertoActionUsers(actionUsers, userIDstring)
	actionToIncrement = actionToIncrement + 1

	return actionToIncrement, oppositeAction, actionUsers, oppositeActionUsers
}

// checks to see if the like (or dislike) is a repeat action, if it is then returns values unchanged
func actionAlreadyPerformed(userIDstring string, actionToIncrement int, oppositeAction int, actionUsers string, oppositeActionUsers string) bool {
	splitActionUsers := strings.Split(actionUsers, ",")
	for _, idAccount := range splitActionUsers {
		if idAccount == userIDstring {
			return true
		}
	}
	return false
}

// checks to see if the oppsite action has already taken place, if so then removes it
// and indicates which user to remove from the oppositeActionUsers list
func oppositeActionChecker(userIDstring string, oppositeActionUsers string) (bool, int) {
	splitOppositeActionUsers := strings.Split(oppositeActionUsers, ",")
	for i, userNumber := range splitOppositeActionUsers {
		if userNumber == userIDstring {
			return true, i
		}

	}
	return false, 0
}

// removes user from oppositeActionUsers list
func removeUserIDFromOppositeActionUsers(oppositeActionUsers string, user int) string {
	splitOppositeActionUsers := strings.Split(oppositeActionUsers, ",")
	splitOppositeActionUsers = append(splitOppositeActionUsers[:user], splitOppositeActionUsers[user+1:]...)
	oppositeActionUsers = strings.Join(splitOppositeActionUsers, ",")
	oppositeActionUsers = strings.TrimPrefix(oppositeActionUsers, ",")
	return oppositeActionUsers
}

// adds userID to actionUsers
func addUsertoActionUsers(actionUsers string, userIDstring string) string {
	actionUsers = actionUsers + "," + userIDstring
	actionUsers = strings.TrimPrefix(actionUsers, ",")
	return actionUsers
}
