-- 004_create_comments_table_up.sql
CREATE TABLE IF NOT EXISTS Comments (
    CommentID INTEGER PRIMARY KEY,
    PostID INTEGER,
    UserID TEXT,
    Body TEXT,
    CreationDate DATETIME DEFAULT CURRENT_TIMESTAMP,
    Likes INTEGER,
    Dislikes INTEGER,
    WhoLiked TEXT,
    WhoDisliked TEXT
);
