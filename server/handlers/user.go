package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"usersvc.io/api/v1/server/requests"
	"usersvc.io/api/v1/server/services"
)

// func UserCtx(next http.Handler) http.Handler {
// }

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func ListUsers(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request requests.CreateUserRequest
	err := readJSON(r, &request)
	if err != nil {
		JSONError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	services.CreateUser(request)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

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
