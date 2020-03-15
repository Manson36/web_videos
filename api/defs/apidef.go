package defs

//requests
type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//Data Models
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	Id         string
	VideoId    string
	AuthorName string
	Content    string
}

type SimpleSession struct {
	UserName string //login_name
	TTL      int64
}
