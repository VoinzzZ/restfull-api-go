package router

import (
	"net/http"
	"restfull-api-go/internal/delivery/http/handler"
)

func SetupRouter(userHandler *handler.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/users", userHandler.CreateUser)
	mux.HandleFunc("GET /api/v1/users/{id}", userHandler.GetUserByID)
	mux.HandleFunc("GET /api/v1/users", userHandler.GetAllUsers)
	mux.HandleFunc("PUT /api/v1/users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /api/v1/users/{id}", userHandler.DeleteUser)
	return mux
}
