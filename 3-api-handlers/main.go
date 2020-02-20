package main

import (
	"log"
	"net/http"

	"github.com/west-la-go/simple-crud-api/3-api-handlers/handlers"
)

func main() {
	log.Println("INFO: starting app")

	var uc handlers.UserController

	router := http.NewServeMux() // mux == multiplexer
	router.HandleFunc("/", handlers.Home)
	router.HandleFunc("/users", uc.Users)

	log.Fatal(http.ListenAndServe(":8080", router))
}
