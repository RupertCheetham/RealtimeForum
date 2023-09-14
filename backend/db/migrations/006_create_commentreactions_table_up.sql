-- 006_create_commentreactions_table_up.sql
CREATE TABLE IF NOT EXISTS COMMENTREACTIONS (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Likes INTEGER,
    Dislikes INTEGER,
    WhoLiked TEXT,
    WhoDisliked TEXT
);