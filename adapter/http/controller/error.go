package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Printf("%s\n", message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Message{
		Status:  status,
		Message: message,
	})
}
