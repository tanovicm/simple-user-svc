package services

import (
	"github.com/Kamva/mgm"
	"usersvc.io/api/v1/server/models"
	"usersvc.io/api/v1/server/requests"
)

func CreateUser(req requests.CreateUserRequest) (*models.User, error) {

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Nickname:  req.Nickname,
		Password:  req.Password,
		Email:     req.Email,
		Country:   req.Country,
	}

	err := mgm.Coll(user).Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(req requests.CreateUserRequest) {

	

}

func DeleteUser(req requests.CreateUserRequest) {

}
func ListUsers(req requests.CreateUserRequest) {

}
