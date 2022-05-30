package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeToBody(writer http.ResponseWriter, body interface{}, code int) {

	bytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("could not convert body to bytes: %s\n", err)
	}

	writer.WriteHeader(code)
	if _, err = writer.Write(bytes); err != nil {
		log.Printf("could not write to body: %s\n", err)
	}
}
