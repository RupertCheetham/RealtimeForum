package db

type UserEntry struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type PostEntry struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Img          string `json:"img"`
	Body         string `json:"body"`
	Categories   string `json:"categories"`
	CreationDate string `json:"creationDate"`
	Likes        int    `json:"likes"`
	Dislikes     int    `json:"dislikes"`
	WhoLiked     string `json:"whoLiked"`
	WhoDisliked  string `json:"whoDisliked"`
}

type CommentEntry struct {
	Id           int    `json:"id"`
	ParentPostID int    `json:"parentPostId"`
	Username     string `json:"username"`
	Body         string `json:"body"`
	CreationDate string `json:"creationDate"`
	Likes        int    `json:"likes"`
	Dislikes     int    `json:"dislikes"`
	WhoLiked     string `json:"whoLiked"`
	WhoDisliked  string `json:"whoDisliked"`
}

type PostWithComments struct {
	Id           int            `json:"id"`
	Username     string         `json:"username"`
	Img          string         `json:"img"`
	Body         string         `json:"body"`
	Categories   string         `json:"categories"`
	CreationDate string         `json:"creationDate"`
	Likes        int            `json:"likes"`
	Dislikes     int            `json:"dislikes"`
	WhoLiked     string         `json:"whoLiked"`
	WhoDisliked  string         `json:"whoDisliked"`
	Comments     []CommentEntry `json:"comments"`
}