package model

type Tweet struct {
	ID      string `json:"ID"`
	UserID  string `json:"userID"`
	User    *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewTweet struct {
	UserID  string `json:"userID"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
