package handlers

import (
	"auth/config"
	"auth/repositories"
	"auth/services"
	"net/http"
	"net/http/httptest"
)

func Start() *httptest.Server {
	cfg := config.NewConfig(true)

	userRepository := repositories.NewUserRepositoryMock()
	tokenService := services.NewTokenService(cfg)

	authHandler := NewAuthHandler(cfg)
	userHandler := NewUserHandler(tokenService, userRepository)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/profile", userHandler.GetProfile)

	return httptest.NewServer(mux)
}
