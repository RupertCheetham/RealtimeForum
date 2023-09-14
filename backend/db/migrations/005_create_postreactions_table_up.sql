-- 005_create_postlikes_table_up.sql
CREATE TABLE IF NOT EXISTS POSTREACTIONS (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Likes INTEGER,
    Dislikes INTEGER,
    WhoLiked TEXT,
    WhoDisliked TEXT
);