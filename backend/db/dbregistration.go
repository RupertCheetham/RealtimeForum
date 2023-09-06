package db

import "log"

func AddRegistrationToDatabase(nickname string, age int, gender string, firstName string, lastName string, email string, password string) error {
	_, err := Database.Exec("INSERT INTO Registration (Nickname, Age, Gender, First_name, Last_name, Email, Password) VALUES (?, ?, ?, ?, ?, ?, ?)", nickname, age, gender, firstName, lastName, email, password)
	if err != nil {
		log.Println("Error adding registration to database:", err)
	}
	return err
}

func GetRegistrationFromDatabase() ([]RegistrationEntry, error) {
	rows, err := Database.Query("SELECT Nickname, Age, Gender, First_name, Last_name, Email, Password FROM Registration ORDER BY Id ASC")
	if err != nil {
		log.Println("Error querying registrations from database:", err)
		return nil, err
	}
	defer rows.Close()

	var registrations []RegistrationEntry
	for rows.Next() {
		var entry RegistrationEntry
		err := rows.Scan(&entry.Nickname, &entry.Age, &entry.Gender, &entry.FirstName, &entry.LastName, &entry.Email, &entry.Password)
		if err != nil {
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		registrations = append(registrations, entry)
	}

	return registrations, nil
}
