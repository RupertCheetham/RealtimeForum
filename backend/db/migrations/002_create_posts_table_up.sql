-- 002_create_posts_table_up.sql
CREATE TABLE IF NOT EXISTS Posts (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Nickname TEXT,
    Img TEXT,
    Body TEXT,
    Categories Text,
    CreationDate DATETIME DEFAULT CURRENT_TIMESTAMP,
    Likes INTEGER,
    Dislikes INTEGER,
    WhoLiked TEXT,
    WhoDisliked TEXT
);
