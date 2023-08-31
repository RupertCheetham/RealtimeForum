package forum

import (
	"fmt"
	"log"
)

// Returns a UserID when given the cookie's value (a string unique to that user)
func cookieUserID(uuid string) string {

	row, err := Database.Query("SELECT * FROM Cookies ORDER BY CreationDate")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var SessionID string
		var UserID string
		var CreationDate string
		row.Scan(&SessionID, &UserID, &CreationDate)
		if uuid == SessionID {
			return UserID
		}
	}
	return "No user found."
}

// Checks to see if User has multiple sessions running (they're only allowed 1)
// and creates a list of extra sessions to be deleted in the next fuction - trying to delete at the same time caused the
// database to lock up
func UserAlreadyLoggedIn(CurrentUserID string) []string {
	var extraSessionIDs []string
	counter := 0

	row, err := Database.Query("SELECT SessionID, UserID FROM Cookies ORDER BY CreationDate DESC")
	if err != nil {
		log.Fatal(err, " In UserAlreadyLoggedIn")
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var SessionID string
		var UserID string
		row.Scan(&SessionID, &UserID)
		if CurrentUserID == UserID {
			counter++
			if counter > 1 {

				extraSessionIDs = append(extraSessionIDs, SessionID)

			}
		}
	}
	return extraSessionIDs
}

// removes any extra user sessions
func ExtraSessionRemover(extraSessionIDs []string) {

	if len(extraSessionIDs) > 0 {

		for i := 0; i < len(extraSessionIDs); i++ {

			_, err := Database.Exec("DELETE FROM Cookies WHERE SessionID = ?", extraSessionIDs[i])
			if err != nil {
				log.Fatal(err, ", in UserAlreadyLoggenIn")
			}
			fmt.Println("Removed a session", extraSessionIDs[i])

		}

	}

}

// checks if a user session is now missing (could deleted by UserAlreadyLoggedIn, or affected by tampering of the cookie)
func UserSessionMissing(uuid string) bool {

	row, err := Database.Query("SELECT SessionID FROM Cookies ORDER BY CreationDate DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var SessionID string
		row.Scan(&SessionID)
		if uuid == SessionID {
			log.Println("The User Still Has A Session")
			return false
		}
	}

	return true
}
