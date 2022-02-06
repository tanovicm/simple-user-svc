package models

import "github.com/Kamva/mgm"

// User contains user
type User struct {
	// MongoDB default model
	mgm.DefaultModel `bson:",inline"`
	// First name of a user
	FirstName string `json:"firstname" bson:"firstname"`
	// Last name of a user
	LastName string `json:"lastname" bson:"lastname"`
	// Nickname of a user
	Nickname string `json:"nickname" bson:"nickname"`
	// Password of a user
	Password string `json:"password" bson:"password"`
	// Email of a user
	Email string `json:"email" bson:"email"`
	// Country of a user
	Country string `json:"country" bson:"country"`
}
