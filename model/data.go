package models

type Data struct {
	Posts       []Post
	Categories  []Category
	Users       []User
	CurrentUser User
	IsConnected bool
}

type DataProfil struct {
	User        User
	Posts       []Post
	IsMe        bool
	IsConnected bool
	CurrentUser User
}

type DataNotFound struct {
	IsConnected bool
	CurrentUser User
}

type DataEditPost struct {
	CurrentUser User
	Post        Post
}

type DataEditCom struct {
	Comment     Comment
	CurrentUser User
	Id          string
}

type DataComment struct {
	IdPost      int
	CurrentUser User
	Lescoms     []Comment
	Post        Post
	Categories  []Category
	IsConnected bool
}

type DataTopic struct {
	Posts       []Post
	Current     Category
	Categories  []Category
	IsConnected bool
	CurrentUser User
}