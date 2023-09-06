-- 003_create_registration_table_up.sql
CREATE TABLE IF NOT EXISTS REGISTRATION (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Username TEXT,
		Age INTEGER,
		Gender TEXT,
		First_name Text,
		Last_name TEXT,
		Email TEXT,
		Password TEXT
	);