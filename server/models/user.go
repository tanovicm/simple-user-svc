package models

import "github.com/Kamva/mgm"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"first_name bson:"first_name"`
	LastName         string `json:"last_name bson:"last_name"`
	Nickname         string `json:"nickname bson:"nickname"`
	Password         string `json:"password bson:"password"`
	Email            string `json:"email bson:"email"`
	Country          string `json:"country bson:"country"`
}
