package models

import "time"

type Url struct {
	LongUrl   string    `json:"long_url" bson:"long_url"`
	ShortUrl  string    `json:"short_url" bson:"short_url"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdateAt  time.Time `json:"update_at" bson:"update_at"`
}

func NewUrl(longUrl string, shortUrl string) Url {
	url := Url{}
	url.LongUrl = longUrl
	url.ShortUrl = shortUrl
	url.CreatedAt = time.Now()
	url.UpdateAt = time.Now()
	return url
}
