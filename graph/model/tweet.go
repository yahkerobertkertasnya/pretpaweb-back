package model

import (
	"github.com/99designs/gqlgen/graphql"
	"time"
)

type Tweet struct {
	ID           string       `json:"ID"`
	UserID       string       `json:"userID"`
	User         *User        `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Content      string       `json:"content"`
	Image        string       `json:"image"`
	CreatedAt    time.Time    `json:"createdAt"`
	Liked        bool         `json:"liked"`
	LikeCount    int          `json:"likeCount"`
	ParentID     *string      `json:"parentID"`
	Parent       *Tweet       `json:"parent,omitempty" gorm:"foreignKey:ParentID;references:ID"`
	Comments     []*Tweet     `json:"comments,omitempty" gorm:"foreignKey:ParentID;references:ID"`
	CommentCount *int         `json:"commentCount,omitempty"`
	Like         []*TweetLike `json:"like,omitempty"  gorm:"foreignKey:TweetID;references:ID"`
}

type TweetLike struct {
	TweetID string `json:"tweetID" gorm:"primaryKey"`
	UserID  string `json:"userID" gorm:"primaryKey"`
	User    *User  `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type NewTweet struct {
	Content  string          `json:"content"`
	ParentID *string         `json:"parentID,omitempty"`
	Image    *graphql.Upload `json:"image,omitempty"`
}
