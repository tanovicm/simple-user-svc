package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
