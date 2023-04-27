package models

import "time"

type Session struct {
	Id             int       `json:"-"`
	UserId         int       `json:"userId"`
	Token          string    `json:"token"`
	ExpirationDate time.Time `json:"expirationDate"`
}
