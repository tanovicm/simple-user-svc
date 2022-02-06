package response

type CreateUserResponse struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Nickname  string `json:"nickname" bson:"nickname"`
	Email     string `json:"email" bson:"email"`
	Country   string `json:"country" bson:"country"`
}
