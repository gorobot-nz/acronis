package apimodels

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint64 `json:"expires_in"`
	ExpiresOn   uint64 `json:"expires_on"`
	IdToken     string `json:"id_token"`
	Scope       string `json:"scope"`
}
