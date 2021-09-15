package helpers

import (
	"auth/helper"
	"auth/responses"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func AssertTokenResponse(t *testing.T, testCase TestCaseValidate, gotClaims *helper.JwtCustomClaims, err error) {
	t.Helper()

	if testCase.WantError {
		assert.Error(t, err)
		assert.Contains(t, err.Error(), testCase.WantErrorMsg)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, testCase.WantID, gotClaims.ID)
	}
}

func AssertUserProfileResponse(t *testing.T, recorder *httptest.ResponseRecorder) {
	t.Helper()

	var response responses.UserResponse
	err := json.Unmarshal([]byte(recorder.Body.String()), &response)

	if assert.NoError(t, err) {
		assert.Equal(t, responses.UserResponse{
			ID:    1,
			Email: "test-1@example.com",
			Name:  "Test User 1",
		}, response)
	}
}
