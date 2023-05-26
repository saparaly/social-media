package models

type User struct {
	Id          int    `json:"id"`
	Img         string `json:"avaimg"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	Followers   []int  `json:"followers"`
	Following   []int  `json:"following"`
	IsFollowing bool   `json:"isFollowing"`
}
