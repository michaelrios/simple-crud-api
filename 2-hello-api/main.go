package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("INFO: starting app")

	router := http.NewServeMux() // mux == multiplexer

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := writer.Write([]byte("hello api"))
		if err != nil {
			log.Println("ERR: failed to write bytes:", err.Error())
			return
		}

		log.Println("INFO: wrote", bytes, "bytes")
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
