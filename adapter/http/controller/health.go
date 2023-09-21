package controller

import (
	"encoding/json"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	type HealthCheckResponse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	health := &HealthCheckResponse{
		Status:  200,
		Message: "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(health)
}
