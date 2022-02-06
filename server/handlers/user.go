package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"usersvc.io/api/v1/server/models"
	"usersvc.io/api/v1/server/requests"
	"usersvc.io/api/v1/server/response"
	"usersvc.io/api/v1/server/services"
)

func UserCtx(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")

		user, err := services.GetUser(userID)
		if err != nil {
			log.Printf("error retrieving user with userID: %v, err: %v\n", userID, err)
			JSONError(w, "error retrieving user", http.StatusInternalServerError)
			return
		}
		if user == nil {
			log.Printf("user with userID: %v not found", userID)
			JSONError(w, "user not found", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(*models.User)

	JSONOk(w, &response.GetUserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Country:   user.Country,
	})
}

func ListUsers(w http.ResponseWriter, r *http.Request) {

	filter := map[string]string{}
	if r.URL.Query().Get("country") != "" {
		filter["country"] = r.URL.Query().Get("country")
	}

	var offset, limit int
	var err error
	if r.URL.Query().Get("limit") != "" && r.URL.Query().Get("offset") != "" {
		offset, err = strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			log.Println("error converting offset to int")
			return
		}
		limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			log.Println("error converting limit to int")
			return
		}
	}

	users, err := services.ListUsers(filter)
	if err != nil {
		log.Printf("error retrieving users err: %v\n", err)
		JSONError(w, "User listing failed", http.StatusInternalServerError)
		return
	}

	resp := []*response.GetUserResponse{}
	for _, user := range users {
		rsp := &response.GetUserResponse{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Nickname:  user.Nickname,
			Email:     user.Email,
			Country:   user.Country,
		}
		resp = append(resp, rsp)
	}
	if offset == 0 && limit == 0 {
		JSONOk(w, resp)
		return
	}

	end := offset + limit
	if end > len(resp) {
		end = len(resp)
	}

	JSONOk(w, resp[offset:end])
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var request requests.CreateUserRequest
	err := readJSON(r, &request)
	if err != nil {
		JSONError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := services.CreateUser(request)
	if err != nil {
		log.Printf("error creating user: err: %v\n", err)
		JSONError(w, "User creation failed", http.StatusInternalServerError)
		return
	}

	JSONOk(w, &response.CreateUserResponse{ID: user.ID.Hex()})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(*models.User)

	var request requests.UpdateUserRequest
	err := readJSON(r, &request)
	if err != nil {
		JSONError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.UpdateUser(user.ID.Hex(), &request)
	if err != nil {
		log.Printf("error updating user with id: %v: err: %v\n", user.ID.Hex(), err)
		JSONError(w, "User update failed", http.StatusInternalServerError)
		return
	}
	JSONOk(w, &struct{}{})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(*models.User)

	err := services.DeleteUser(user.ID.Hex())
	if err != nil {
		log.Printf("error deleting user with id: %v: err: %v\n", user.ID.Hex(), err.Error())
		JSONError(w, "User deletion failed", http.StatusInternalServerError)
		return
	}

	JSONOk(w, &struct{}{})
}
