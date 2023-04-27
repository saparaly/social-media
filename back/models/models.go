package models

import "errors"

type Error struct {
	Message string `json:"message"`
}

var ErrNoRecord = errors.New("models: no matching record found")
