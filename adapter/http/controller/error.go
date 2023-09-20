package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	log.Printf("%s\n", message)

	type ErrorMessage struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorMessage{
		Status:  statusCode,
		Message: message,
	})
}
