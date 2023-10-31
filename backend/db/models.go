package db

import "time"

type UserEntry struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Session struct {
	SessionID      string
	UserId         int
	CreationTime   time.Time
	ExpirationTime time.Time
}

type PostEntry struct {
	Id           int      `json:"id"`
	UserId       int      `json:"userID"`
	Username     string   `json:"username"`
	Img          string   `json:"img"`
	Body         string   `json:"body"`
	Categories   []string `json:"categories"`
	CreationDate string   `json:"creationDate"`
	ReactionID   int      `json:"reactionID"`
	Likes        int      `json:"postLikes"`
	Dislikes     int      `json:"postDislikes"`
}

type CommentEntry struct {
	Id           int    `json:"id"`
	ParentPostID int    `json:"parentPostId"`
	UserId       int    `json:"userID"`
	Username     string `json:"username"`
	Body         string `json:"body"`
	CreationDate string `json:"creationDate"`
	ReactionID   int    `json:"reactionID"`
	Likes        int    `json:"commentLikes"`
	Dislikes     int    `json:"commentDislikes"`
}

type ReactionEntry struct {
	UserID     int    `json:"userID"`
	Type       string `json:"type"`
	ParentID   int    `json:"parentID"`
	Action     string `json:"action"`
	ReactionID int    `json:"reactionID"`
}

type Reaction struct {
	Id          int    `json:"id"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	WhoLiked    string `json:"wholiked"`
	WhoDisliked string `json:"whodisliked"`
}

type ChatMessage struct {
	Type      string `json:"type"`
	Body      string `json:"body"`
	Sender    int    `json:"sender"`
	Recipient int    `json:"recipient"`
	Time      string `json:"time"`
}

type ChatInfo struct {
	RecentChat   []UserEntry `json:"recentChat"`
	Alphabetical []UserEntry `json:"alphabetical"`
}
