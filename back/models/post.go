package models

import (
	"time"
)

type Post struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	Title       string `json:"title"`
	Img         string `json:"img"`
	Description string `json:"description"`
	// TagsId      []int      `json:"tags"`
	Tags         []string   `json:"tags"`
	Created      *time.Time `json:"created"`
	Date         string     `json:"date"`
	Location     string     `json:"location"`
	Like         int        `json:"like"`
	Dislike      int        `json:"dislike"`
	ReactionType string     `json:"ReactionType"`
	/**/
	Username string `json:"postusername"`
	UserRole string `json:"postuserrole"`
	IsUser   bool   `json:"trueuser"`
	Reacted  bool   `json:"reactedtopost"`
}

type Like struct {
	Id        int `json:"id"`
	PostId    int `json:"postid"`
	CommentId int `json:"commentid"`
	UserId    int `json:"userid"`
}

type Dislike struct {
	Id        int `json:"id"`
	PostId    int `json:"postid"`
	CommentId int `json:"commentid"`
	UserId    int `json:"userid"`
}

type Comment struct {
	Id           int    `json:"id"`
	UserId       int    `json:"userid"`
	PostId       int    `json:"postid"`
	Text         string `json:"text"`
	Like         int    `json:"like"`
	Dislike      int    `json:"dislike"`
	ReactionType string `json:"ReactionType"`
	/**/
	Username string `json:"commentusername"`
	UserRole string `json:"commentuserrole"`
	Reacted  bool   `json:"reacted"`
}
