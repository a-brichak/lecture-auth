package main

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strings"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func ValidateBearerToken(authHeader string) (*JwtCustomClaims, error) {
	tokenString := BearerAuthHeader(authHeader)

	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}

func BearerAuthHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}
