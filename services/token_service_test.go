package services

import (
	"auth/config"
	"auth/helper"
	"auth/tests/helpers"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetTokenFromBearerString(t *testing.T) {
	s := NewTokenService(&config.Config{})
	testCases := []helpers.TestCaseGetBearerToken{
		{
			BearerString: "Bearer test_token",
			Want: "test_token",
		},
		{
			BearerString: "Beare test_token",
			Want: "",
		},
	}

	for _, testCase := range testCases{
		got := s.GetTokenFromBearerString(testCase.BearerString)
		assert.Equal(t, testCase.Want, got)
	}
}

func TestGenerateAccessToken(t *testing.T) {
	userID := 1
	cfg := &config.Config{
		AccessSecret:           "access",
		AccessLifetimeMinutes:  1,
	}
	s := NewTokenService(cfg)

	tokenString, err := s.GenerateAccessToken(userID)

	assert.NoError(t, err)

	token, err := jwt.ParseWithClaims(tokenString, &helper.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AccessSecret), nil
	})
	assert.NoError(t, err)

	claims, ok := token.Claims.(*helper.JwtCustomClaims)
	assert.True(t, ok)
	assert.True(t, token.Valid)

	got := claims.ID
	assert.Equal(t, userID, got)
}

func TestValidateAccessToken(t *testing.T) {
	userID := 1
	cfg := &config.Config{
		AccessSecret:           "access",
		AccessLifetimeMinutes:  1,
		RefreshSecret:           "refresh",
		RefreshLifetimeMinutes: 1,
	}
	s := NewTokenService(cfg)
	tokenString, _ := s.GenerateAccessToken(userID)
	refreshTokenString, _ := s.GenerateRefreshToken(userID)
	invalidTokenString := tokenString + "a"

	cfg.AccessLifetimeMinutes = 0
	expiredTokenString, _ := s.GenerateAccessToken(userID)

	testCases := []helpers.TestCaseValidate{
		{
			Name: "Valid token",
			AccessToken: tokenString,
			WantError: false,
			WantErrorMsg: "",
			WantID: userID,
		},
		{
			Name: "Invalid token",
			AccessToken: invalidTokenString,
			WantError: true,
			WantErrorMsg: "signature is invalid",
			WantID: 0,
		},
		{
			Name: "Token with non-expected signature",
			AccessToken: refreshTokenString,
			WantError: true,
			WantErrorMsg: "signature is invalid",
			WantID: 0,
		},
		{
			Name: "Expired token",
			AccessToken: expiredTokenString,
			WantError: true,
			WantErrorMsg: "token is expired",
			WantID: 0,
		},
	}

	for _, testCase := range testCases{
		t.Run(testCase.Name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotClaims, err := s.ValidateAccessToken(testCase.AccessToken)

			helpers.AssertTokenResponse(t, testCase, gotClaims, err)
		})
	}
}
