-- 003_create_registration_table_up.sql
CREATE TABLE IF NOT EXISTS Registration (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Nickname TEXT,
		Age INTEGER,
		Gender TEXT,
		First_name Text,
		Last_name TEXT,
		Email TEXT,
		Password TEXT
	);