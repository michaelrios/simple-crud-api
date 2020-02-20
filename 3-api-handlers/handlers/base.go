package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Respond(w http.ResponseWriter, data interface{}, statusCode int) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(bytes)

	return err
}

func Home(writer http.ResponseWriter, request *http.Request) {
	bytes, err := writer.Write([]byte("welcome home!"))
	if err != nil {
		log.Println("ERR: failed to write bytes:", err.Error())
		return
	}

	log.Println("INFO: wrote", bytes, "bytes")
}
