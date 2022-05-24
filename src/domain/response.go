package domain

type ResponseWithMessage struct {
	Message string `json:"message"`
}

type ResponseWithMessageAndResults struct {
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type CognitoTokens struct {
	AccessToken  string `json:"accessToken"`
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}
