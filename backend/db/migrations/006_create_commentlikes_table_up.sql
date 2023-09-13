-- 006_create_commentlikes_table_up.sql
CREATE TABLE IF NOT EXISTS COMMENTLIKES (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Likes INTEGER,
    Dislikes INTEGER,
    WhoLiked TEXT,
    WhoDisliked TEXT
);