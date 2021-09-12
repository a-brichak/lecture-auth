package helpers

import (
	"auth/helper"
	"github.com/stretchr/testify/assert"
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
