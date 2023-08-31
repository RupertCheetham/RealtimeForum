package forum

type Activity struct {
	LoggedInUsername    string
	MyPosts             []Post
	MyLikedPosts        []Post
	NoPostsMessage      string
	NoLikedPostsMessage string
}

type Authentication struct {
	Email string
	Name  string
}

type Category struct {
	LoggedInUsername string
	Category         string
	Posts            []Post
	NoPostsMessage   string
}

type Comment struct {
	CommentID    int
	PostID       int
	Username     string
	Body         string
	CreationDate string
	Likes        int
	Dislikes     int
	whoLiked     string
	whoDisliked  string
}

type ErrorMessage struct {
	UsernameError  string
	EmailError     string
	PasswordError  string
	ConfirmError   string
	AlreadySigneUp string
	SignInError    string
}

type Post struct {
	PostID       int
	Username     string
	Img          string
	Body         string
	CreationDate string
	Categories   []string
	Likes        int
	Dislikes     int
	whoLiked     string
	whoDisliked  string
	Comments     []Comment
	CommentCount int
}

type FinalStruct struct {
	LoggedInUsername string
	Posts            []Post
}
