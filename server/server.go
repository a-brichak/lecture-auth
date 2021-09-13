package server

import (
	"auth/config"
	"auth/repositories"
	"auth/server/handlers"
	"auth/services"
	"log"
	"net/http"
)

func Start(cfg *config.Config) {
	userRepository := repositories.NewUserRepository()
	tokenService := services.NewTokenService(cfg)

	authHandler := handlers.NewAuthHandler(cfg)
	userHandler := handlers.NewUserHandler(tokenService, userRepository)

	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/profile", userHandler.GetProfile)

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
