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

// UsersList contains a list of orders
type UsersList struct {
	// A list of orders
	Users []*User `json:"items"`
	// The id to query the next page
	NextPageID int `json:"next_page_id,omitempty" example:"10"`
} // @name UsersList
