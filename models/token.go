package models

import "time"

type Token struct {
	Token     string    `json:"token" bson:"token"`
	Username  string    `json:"username" bson:"username"` // Reference to User
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	ExpireAt  time.Time `json:"expire_at" bson:"expire_at"`
}

func NewToken(token string, username string) Token {
	t := Token{}
	t.Token = token
	t.Username = username
	t.CreatedAt = time.Now()
	t.ExpireAt = time.Now().Add(time.Hour * 24 * 7) // 7 days to expire
	return t
}
