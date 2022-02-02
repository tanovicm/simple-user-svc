package server

import (
	"log"
	"net/http"
)

func StartServer() error {

	mux := BuildRouter()

	log.Println("Listening on port 8000")
	return http.ListenAndServe(":8000", mux)

}
