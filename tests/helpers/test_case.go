package helpers

type TestCaseGetBearerToken struct {
	BearerString string
	Want string
}

type TestCaseValidate struct {
	Name string
	AccessToken string
	WantError bool
	WantErrorMsg string
	WantID int
}
