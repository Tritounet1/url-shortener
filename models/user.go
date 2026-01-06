package models

import "time"

type User struct {
	Email     string    `json:"email" bson:"email"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdateAt  time.Time `json:"update_at" bson:"update_at"`
}

func NewUser(email string, username string, password string) User {
	user := User{}
	user.Email = email
	user.Username = username
	user.Password = password
	user.CreatedAt = time.Now()
	user.UpdateAt = time.Now()
	return user
}
