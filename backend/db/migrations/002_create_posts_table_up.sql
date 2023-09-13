-- 002_create_posts_table_up.sql
CREATE TABLE IF NOT EXISTS POSTS (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    FOREIGN KEY (Username) REFERENCES USERS(Username),
    Img TEXT,
    Body TEXT,
    Categories Text,
    CreationDate DATETIME DEFAULT CURRENT_TIMESTAMP,
    Likes INTEGER,
    Dislikes INTEGER,
    WhoLiked TEXT,
    WhoDisliked TEXT
);
