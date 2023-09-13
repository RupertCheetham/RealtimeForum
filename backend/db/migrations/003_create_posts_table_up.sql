-- 003_create_posts_table_up.sql
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
    WhoDisliked TEXT,
    FOREIGN KEY (Username) REFERENCES USERS(Username),
    FOREIGN KEY (Likes) REFERENCES POSTLIKES(Likes),
    FOREIGN KEY (Dislikes) REFERENCES POSTLIKES(Dislikes),
    FOREIGN KEY (WhoLiked) REFERENCES POSTLIKES(WhoLiked),
    FOREIGN KEY (WhoDisliked) REFERENCES POSTLIKES(WhoDisliked)
);
