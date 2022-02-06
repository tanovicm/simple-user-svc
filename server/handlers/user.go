package handlers

import (
	"context"
	"encoding/json"
	"fmt"
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
			return
		}
		if user == nil {
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
			fmt.Println("error converting offset to int")
			return
		}
		limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			fmt.Println("error converting limit to int")
			return
		}
	}

	users, err := services.ListUsers(filter)
	if err != nil {
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
	if offset == 0 || limit == 0 {
		JSONOk(w, resp)
		return
	}

	JSONOk(w, resp[offset:offset+limit])

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
		JSONError(w, "Error creating user", http.StatusInternalServerError)
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
		return
	}
	JSONOk(w, &struct{}{})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(*models.User)

	err := services.DeleteUser(user.ID.Hex())
	if err != nil {
		return
	}

	JSONOk(w, &struct{}{})
}

func readJSON(r *http.Request, v interface{}) error {

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return fmt.Errorf("invalid JSON input: %v", err)
	}

	return nil
}

type JSONErr struct {
	Err string `json:"err"`
}

func JSONError(w http.ResponseWriter, errStr string, code int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	writeJSON(w, &JSONErr{Err: errStr})
}

func writeJSON(w http.ResponseWriter, v interface{}) {

	b, err := json.Marshal(v)
	if err != nil {
		http.Error(w, fmt.Sprintf("json encoding error: %v", err), http.StatusInternalServerError)
		return
	}

	writeBytes(w, b)
}

func writeBytes(w http.ResponseWriter, b []byte) {

	_, err := w.Write(b)
	if err != nil {
		http.Error(w, fmt.Sprintf("write error: %v", err), http.StatusInternalServerError)
		return
	}
}

func JSONOk(w http.ResponseWriter, v interface{}) {

	w.Header().Set("Content-Type", "application/json")
	writeJSON(w, v)
}
