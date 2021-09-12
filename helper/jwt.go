package helper

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
