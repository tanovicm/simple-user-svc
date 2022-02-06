package server

import (
	"fmt"
	"log"
	"net/http"

	"usersvc.io/api/v1/lib"
)

func StartServer() error {

	err := lib.OpenDB()
	if err != nil {
		return fmt.Errorf("open db %v", err)
	}
	mux := BuildRouter()

	log.Println("Listening on port 8000")
	return http.ListenAndServe(":8000", mux)

}
