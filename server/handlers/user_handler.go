package handlers

import (
	"auth/repositories"
	"auth/responses"
	"auth/services"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	tokenService   *services.TokenService
	userRepository repositories.IUserRepository
}

func NewUserHandler(tokenService *services.TokenService, userRepository repositories.IUserRepository) *UserHandler {
	return &UserHandler{
		tokenService:   tokenService,
		userRepository: userRepository,
	}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		requestToken := h.tokenService.GetTokenFromBearerString(r.Header.Get("Authorization"))
		claims, err := h.tokenService.ValidateAccessToken(requestToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := h.userRepository.GetUserByID(claims.ID)
		if err != nil {
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		}

		resp := responses.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
