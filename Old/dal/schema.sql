CREATE TABLE IF NOT EXISTS Users (
    UserID INTEGER PRIMARY KEY,
    Username TEXT NOT NULL,
    Email TEXT,
    Password TEXT,
    RegistrationDate DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Posts (
    PostID INTEGER PRIMARY KEY,
    UserID TEXT,
    Img TEXT,
    Body TEXT,
    Categories TEXT,
    CreationDate  DATETIME DEFAULT CURRENT_TIMESTAMP,
    Likes INTEGER,
    Dislikes INTEGER,
    WhoLiked TEXT,
    WhoDisliked TEXT
);

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

CREATE TABLE IF NOT EXISTS Cookies (
    SessionID TEXT NULL,
    UserID TEXT NOT NULL,
    CreationDate DATETIME DEFAULT CURRENT_TIMESTAMP
);