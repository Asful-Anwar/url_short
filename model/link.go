package model

import "time"

type Link struct {
	ID     int
	Link   string
	Newlink string
	UserID int
	Title  string
	Description string
	CreatedAt time.Time
	ModifiedAt time.Time
	ExpiredAt time.Time
}