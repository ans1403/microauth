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
