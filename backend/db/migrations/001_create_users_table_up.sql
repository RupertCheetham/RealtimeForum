-- 001_create_users_table_up.sql
CREATE TABLE IF NOT EXISTS Users (
    UserID INTEGER PRIMARY KEY,
    Username TEXT NOT NULL,
    Email TEXT,
    Password TEXT,
    RegistrationDate DATETIME DEFAULT CURRENT_TIMESTAMP
);
