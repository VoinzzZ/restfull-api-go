package router

import (
	"net/http"
	"restfull-api-go/internal/delivery/http/handler"
)

func registerUserRoutes(mux *http.ServeMux, h *handler.UserHandler) {
	const prefix = "/api/v1/users"

	mux.HandleFunc("GET "+prefix, h.GetAllUsers)
	mux.HandleFunc("POST "+prefix, h.CreateUser)
	mux.HandleFunc("GET "+prefix+"/{id}", h.GetUserByID)
	mux.HandleFunc("PUT "+prefix+"/{id}", h.UpdateUser)
	mux.HandleFunc("DELETE "+prefix+"/{id}", h.DeleteUser)
}
