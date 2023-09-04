package main

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
