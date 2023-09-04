package main

type RegistrationEntry struct {
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type PostEntry struct {
	Id           int    `json:"Id"`
	Nickname     string `json:"Nickname"`
	Img          string `json:"Img"`
	Body         string `json:"Body"`
	Categories   string `json:"Categories"`
	CreationDate string `json:"CreationDate"`
	Likes        int    `json:"Likes"`
	Dislikes     int    `json:"Dislikes"`
	WhoLiked     string `json:"WhoLiked"`
	WhoDisliked  string `json:"WhoDisliked"`
}
