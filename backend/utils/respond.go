package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// can you write for me the Message func here
func Message(msg string) map[string]interface{} {
	return map[string]interface{}{"Error!!!": msg}
}

// can you write for me the Respond func here
func Respond(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}