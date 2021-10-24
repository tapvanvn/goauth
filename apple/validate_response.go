package apple

type ValidateResponse struct {
	AccessToken string `json:"access_token"`

	TokenType string `json:"token_type"`

	ExpiresIn int `json:"expires_in"`

	RefreshToken string `json:"refresh_token"`

	IDToken string `json:"id_token"`

	Error string `json:"error"`

	ErrorDescription string `json:"error_description"`
}
