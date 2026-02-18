package router

import (
	"net/http"
	"restfull-api-go/internal/delivery/http/handler"
	"restfull-api-go/internal/delivery/http/middleware"
)

func SetupRouter(userHandler *handler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	registerUserRoutes(mux, userHandler)

	var h http.Handler = mux
	h = middleware.Logger(h)

	return h
}
