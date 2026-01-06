package models

import "time"

type Visitor struct {
	IpAdress  string    `json:"ip_adress" bson:"ip_adress"`
	NbClicks  int       `json:"nb_clicks" bson:"nb_clicks"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdateAt  time.Time `json:"update_at" bson:"update_at"`
}

func NewVisitor(ipAdress string) Visitor {
	visitor := Visitor{}
	visitor.CreatedAt = time.Now()
	visitor.UpdateAt = time.Now()
	return visitor
}
