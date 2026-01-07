package models

import "time"

type User struct {
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdateAt  time.Time `json:"update_at" bson:"update_at"`
}

func NewUser(username string, password string) User {
	user := User{}
	user.Username = username
	user.Password = password
	user.CreatedAt = time.Now()
	user.UpdateAt = time.Now()
	return user
}
