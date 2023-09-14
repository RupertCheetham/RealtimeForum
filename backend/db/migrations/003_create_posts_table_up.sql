-- 003_create_posts_table_up.sql
CREATE TABLE IF NOT EXISTS POSTS (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    UserID INTEGER,
    Img TEXT,
    Body TEXT,
    Categories Text,
    CreationDate DATETIME DEFAULT CURRENT_TIMESTAMP,
    ReactionID INTEGER DEFAULT 0,
    FOREIGN KEY (UserID) REFERENCES USERS(Id) ON DELETE SET NULL,
    FOREIGN KEY (ReactionID) REFERENCES POSTREACTIONS(Id) ON DELETE SET NULL
);
