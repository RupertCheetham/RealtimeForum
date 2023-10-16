package db

import (
	"database/sql"
	"realtimeForum/utils"
	"time"
)

// creates a session
func CreateSession(userId int) (session Session, err error) {
	statement := `INSERT INTO Cookies (SessionID, UserID, CreationDate) values (?, ?, ?) returning SessionID, UserID, CreationDate`

	stmt, err := Database.Prepare(statement)
	utils.HandleError("session error:", err)

	defer stmt.Close()

	UUID := utils.GenerateNewUUID()
	timeNow := time.Now()

	err = stmt.QueryRow(UUID, userId, timeNow).Scan(&session.SessionID, &session.UserId, &session.CreationDate)
	return
}

// GetSessionByToken retrieves a session by its session token from the database.
func GetSessionByToken(sessionToken string) (*Session, error) {
	var session Session

	query := "SELECT UserID, ExpirationTime FROM sessions WHERE SessionToken = ?"
	err := Database.QueryRow(query, sessionToken).Scan(&session.UserId, &session.ExpirationTime)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Session not found
		}
		return nil, err // Other error
	}

	session.SessionID = sessionToken

	return &session, nil
}