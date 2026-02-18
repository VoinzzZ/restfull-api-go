package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseJSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
}

func ResponseError(w http.ResponseWriter, statusCode int, message string) {
	ResponseJSON(w, statusCode, message, nil)
}
