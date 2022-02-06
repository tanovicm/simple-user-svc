package services

import (
	"testing"

	"usersvc.io/api/v1/lib"
	"usersvc.io/api/v1/server/models"
	"usersvc.io/api/v1/server/requests"
)

func TestCreateUser(t *testing.T) {

	_ = lib.OpenDB()

	type args struct {
		req requests.CreateUserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{name: "pass", args: args{requests.CreateUserRequest{
			FirstName: "Marijana",
			LastName:  "Tanovic",
			Nickname:  "Mara",
			Password:  "123",
			Email:     "mail@mail.com",
			Country:   "Serbia",
		}}, want: &models.User{

			FirstName: "Marijana",
			LastName:  "Tanovic",
			Nickname:  "Mara",
			Password:  "123",
			Email:     "mail@mail.com",
			Country:   "Serbia",
		}, wantErr: false},
		{name: "fail", args: args{requests.CreateUserRequest{
			FirstName: "Marijana",
			LastName:  "Tanovic",
			Nickname:  "Mara",
			Password:  "ssss",
			Email:     "mail@mail.com",
			Country:   "Serbia",
		}}, want: &models.User{

			FirstName: "Marijana",
			LastName:  "Tanovic",
			Nickname:  "Mara",
			Password:  "123",
			Email:     "mail@mail.com",
			Country:   "Serbia",
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := CreateUser(tt.args.req)
			if !compareStructs(got, tt.want) && !tt.wantErr {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareStructs(got *models.User, want *models.User) bool {

	if got.FirstName != want.FirstName {
		return false
	}
	if got.LastName != want.LastName {
		return false
	}
	if got.Nickname != want.Nickname {
		return false
	}
	if got.Password != want.Password {
		return false
	}
	if got.Country != want.Country {
		return false
	}
	if got.Email != want.Email {
		return false
	}
	return true
}
