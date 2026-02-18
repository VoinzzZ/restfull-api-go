package main

import (
	"log"
	"net/http"
	"restfull-api-go/config"
	"restfull-api-go/internal/delivery/http/handler"
	"restfull-api-go/internal/delivery/http/router"
	"restfull-api-go/internal/repository"
	"restfull-api-go/internal/usecase"
	"restfull-api-go/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := database.NewMySQLConnection(cfg.GetDSN())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	mux := router.SetupRouter(userHandler)

	serverAddr := ":" + cfg.PORT
	log.Printf("Server starting on http://localhost%s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, mux))
}
