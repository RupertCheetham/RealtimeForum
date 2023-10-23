-- 001_create_users_table_up.sql
CREATE TABLE IF NOT EXISTS USERS (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	Username TEXT,
	Age INTEGER,
	Gender TEXT,
	First_name Text,
	Last_name TEXT,
	Email TEXT,
	Password TEXT,
	UNIQUE (Username, Email)
);