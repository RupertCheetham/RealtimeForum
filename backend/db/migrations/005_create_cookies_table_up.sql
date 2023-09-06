-- 005_create_cookies_table_up.sql
CREATE TABLE IF NOT EXISTS Cookies (
    SessionID TEXT NULL,
    UserID TEXT NOT NULL,
    CreationDate DATETIME DEFAULT CURRENT_TIMESTAMP
);