package handlers

import (
	"auth/config"
	"auth/repositories"
	"auth/services"
	"auth/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

const userID = 1

type UserHandlerTestSuite struct {
	suite.Suite
	accessToken string
	userHandler *UserHandler
	testSrv     *httptest.Server
}

func (suite *UserHandlerTestSuite) SetupSuite() {
	cfg := &config.Config{
		AccessSecret:          "access",
		AccessLifetimeMinutes: 1,
	}
	tokenService := services.NewTokenService(cfg)

	suite.userHandler = NewUserHandler(tokenService, repositories.NewUserRepositoryMock())
	suite.accessToken, _ = tokenService.GenerateAccessToken(userID)

	suite.testSrv = Start()
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) TestWalkUserHandlerGetProfile() {
	t := suite.T()
	handlerFunc := suite.userHandler.GetProfile
	cases := []helpers.TestCaseHandler{
		{
			TestName: "Successfully get user profile",
			Request: helpers.Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: suite.accessToken,
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "test-1@example.com",
			},
		},
		{
			TestName: "Unauthorized getting user profile",
			Request: helpers.Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: "",
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "invalid",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			request, recorder := helpers.PrepareHandlerTestCase(test)
			test.HandlerFunc(recorder, request)

			assert.Contains(t, recorder.Body.String(), test.Want.BodyPart)
			if assert.Equal(t, test.Want.StatusCode, recorder.Code) {
				if recorder.Code == http.StatusOK {
					helpers.AssertUserProfileResponse(t, recorder)
				}
			}
		})
	}
}

func (suite *UserHandlerTestSuite) TestWalkApiGetProfile() {
	t := suite.T()
	cases := []helpers.TestCaseHandler{
		{
			TestName: "Successfully get user profile",
			Request: helpers.Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: suite.accessToken,
			},
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "test-1@example.com",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			resp, err := suite.testSrv.Client().Get(suite.testSrv.URL + test.Request.Url)
			if assert.NoError(t, err) {
				assert.Equal(t, test.Want.StatusCode, resp.StatusCode)
			}
		})
	}
}
