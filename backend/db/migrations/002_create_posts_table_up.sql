-- 002_create_posts_table_up.sql
CREATE TABLE IF NOT EXISTS POSTS (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Username TEXT,
    Img TEXT,
    Body TEXT,
    Categories Text,
    CreationDate DATETIME DEFAULT CURRENT_TIMESTAMP,
    Likes INTEGER,
    Dislikes INTEGER,
    WhoLiked TEXT,
    WhoDisliked TEXT
);
