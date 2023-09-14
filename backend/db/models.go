package db

type UserEntry struct {
	ID        int    `json:"id"`
	Username  string `json:"userID"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type PostEntry struct {
	Id           int    `json:"id"`
	UserId       int    `json:"userID"`
	Img          string `json:"img"`
	Body         string `json:"body"`
	Categories   string `json:"categories"`
	CreationDate string `json:"creationDate"`
	ReactionID   int    `json:"reactionID"`
}

type CommentEntry struct {
	Id           int    `json:"id"`
	ParentPostID int    `json:"parentPostId"`
	UserId       int    `json:"username"`
	Body         string `json:"body"`
	CreationDate string `json:"creationDate"`
	ReactionID   int    `json:"reactionID"`
}

type Reaction struct {
	Id          int    `json:"id"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	WhoLiked    string `json:"wholiked"`
	WhoDisliked string `json:"whodisliked"`
}
