package services

import (
	"context"
	"fmt"

	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetUser(userID string) (*models.User, error) {

	var user models.User
	err := mgm.Coll(&user).FindByID(userID, &user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("db err: ", err)
	}

	return &user, nil
}

func DeleteUser(userID string) error {

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("user id not ok")
	}

	filter := &bson.M{
		"_id": userObjectID,
	}

	res, err := mgm.Coll(&models.User{}).DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting user %v from db", userID)
	}

	if res.DeletedCount < 1 {
		return fmt.Errorf("user not found to be deleted")
	}

	return nil
}
func ListUsers(filters map[string]string) ([]*models.User, error) {

	users := []*models.User{}

	filter := bson.M{}
	for k, v := range filters {
		filter[k] = v
	}

	err := mgm.Coll(&models.User{}).SimpleFind(&users, &filter, &options.FindOptions{Sort: &bson.M{"created_at": -1}})
	if err != nil {
		return nil, fmt.Errorf("error listing users: %v", err)
	}

	return users, nil
}

func UpdateUser(userID string, req *requests.UpdateUserRequest) error {

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("userID invalid")
	}

	update := bson.M{
		"$set": bson.M{
			"firstname": req.FirstName,
			"lastname":  req.LastName,
			"nickname":  req.Nickname,
			"password":  req.Password,
			"email":     req.Email,
			"country":   req.Country,
		},
	}

	res, err := mgm.Coll(&models.User{}).UpdateOne(context.Background(), bson.M{"_id": userObjectID}, update)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	if res.MatchedCount < 1 {
		return fmt.Errorf("user not found")
	}
	return nil
}
