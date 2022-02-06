package models

import "github.com/Kamva/mgm"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"firstname" bson:"firstname"`
	LastName         string `json:"lastname" bson:"lastname"`
	Nickname         string `json:"nickname" bson:"nickname"`
	Password         string `json:"password" bson:"password"`
	Email            string `json:"email" bson:"email"`
	Country          string `json:"country" bson:"country"`
}
