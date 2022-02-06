package services

import (
	"reflect"
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
			_ = DeleteUser(got.ID.Hex())
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

func TestGetUser(t *testing.T) {
	_ = lib.OpenDB()

	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{name: "hex string is not a valid", args: args{"123"}, want: nil, wantErr: true},
		{name: "non-exiting id", args: args{"000000000000000000000000"}, want: nil, wantErr: false},
		{name: "success", args: args{"000000000000000000000000"}, want: nil, wantErr: false},
	}
	for i, tt := range tests {
		var user *models.User
		if i == len(tests) {
			user, _ = CreateUser(requests.CreateUserRequest{
				FirstName: "Marijana",
				LastName:  "Tanovic",
				Nickname:  "Mara",
				Password:  "ssss",
				Email:     "mail@mail.com",
				Country:   "Serbia",
			})
			tests[len(tests)].args.userID = user.ID.Hex()
			tests[len(tests)].want = user

		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUser(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
			if got != nil {
				_ = DeleteUser(got.ID.Hex())
			}
		})
	}

}

func TestDeleteUser(t *testing.T) {
	_ = lib.OpenDB()
	user, _ := CreateUser(requests.CreateUserRequest{
		FirstName: "Marijana",
		LastName:  "Tanovic",
		Nickname:  "Mara",
		Password:  "ssss",
		Email:     "mail@mail.com",
		Country:   "Serbia",
	})
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "hex string is not a valid", args: args{"123"}, wantErr: true},
		{name: "user-not found", args: args{"000000000000000000000000"}, wantErr: true},
		{name: "success", args: args{user.ID.Hex()}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUser(tt.args.userID); (err != nil) != tt.wantErr {
				t.Log(tt.args.userID)
				t.Errorf("DeleteUser() error = %v, wantErr %v, %v", err, tt.wantErr, tt.name)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	_ = lib.OpenDB()
	user, _ := CreateUser(requests.CreateUserRequest{
		FirstName: "Marijana",
		LastName:  "Tanovic",
		Nickname:  "Mara",
		Password:  "ssss",
		Email:     "mail@mail.com",
		Country:   "Serbia",
	})
	type args struct {
		userID string
		req    *requests.UpdateUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "userID invalid", args: args{
			userID: "000",
			req: &requests.UpdateUserRequest{
				FirstName: "",
				LastName:  "",
				Nickname:  "",
				Password:  "",
				Email:     "",
				Country:   "",
			}}, wantErr: true},
		{name: "userID not found", args: args{
			userID: "000000008a90f2e849ek1754",
			req: &requests.UpdateUserRequest{
				FirstName: "",
				LastName:  "",
				Nickname:  "",
				Password:  "",
				Email:     "",
				Country:   "",
			}}, wantErr: true},
		{name: "success", args: args{
			userID: user.ID.Hex(),
			req: &requests.UpdateUserRequest{
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Nickname:  "NewNickname",
				Password:  user.Password,
				Email:     user.Email,
				Country:   user.Country,
			}}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUser(tt.args.userID, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	_ = DeleteUser(user.ID.Hex())
}
