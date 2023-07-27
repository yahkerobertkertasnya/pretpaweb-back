package model

import "time"

type Tweet struct {
	ID        string    `json:"ID"`
	UserID    string    `json:"userID"`
	User      *User     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewTweet struct {
	Content string `json:"content"`
}
