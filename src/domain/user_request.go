package domain

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ConfirmSignUpRequest struct {
	Username         string `json:"username"`
	ConfirmationCode string `json:"confirmationCode"`
}

type ForgotPasswordRequest struct {
	Username string `json:"email"`
}

type ConfirmForgotPasswordRequest struct {
	Username         string `json:"username"`
	ConfirmationCode string `json:"confirmationCode"`
	NewPassword      string `json:"newPassword"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
