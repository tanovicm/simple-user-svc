package main

import (
	"log"

	"usersvc.io/api/v1/server"
)

func main() {

	log.Fatal(server.StartServer())
}
