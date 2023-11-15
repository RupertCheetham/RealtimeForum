package db

import (
	"database/sql"
	"realtimeForum/utils"
	"time"
)

// creates a session
func CreateSession(userId int, expirationTime time.Time) (session Session, err error) {
	statement := `INSERT INTO Cookies (SessionID, UserID, CreationTime, ExpirationTime) values (?, ?, ?, ?) returning SessionID, UserID, CreationTime, ExpirationTime`

	stmt, err := Database.Prepare(statement)
	utils.HandleError("session error:", err)

	defer stmt.Close()

	UUID := utils.GenerateNewUUID()
	timeNow := time.Now()

	err = stmt.QueryRow(UUID, userId, timeNow, expirationTime).Scan(&session.SessionID, &session.UserId, &session.CreationTime, &session.ExpirationTime)
	return
}

// GetSessionByToken retrieves a session by its session token from the database.
func GetSessionByToken(sessionToken string) (*Session, error) {

	var session Session

	statement := "SELECT UserID, ExpirationTime FROM Cookies WHERE SessionID = ?"
	err := Database.QueryRow(statement, sessionToken).Scan(&session.UserId, &session.ExpirationTime)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Session not found
		}
		return nil, err // Other error
	}

	session.SessionID = sessionToken

	return &session, nil
}
