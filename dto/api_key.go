package dto

type CreateAPIKeyRequest struct {
	Name      string `json:"name" binding:"required"`
	ExpiresAt string `json:"expires_at,omitempty"`
	Scopes    string `json:"scopes,omitempty"`
}

type CreateAPIKeyResponse struct {
	APIKey    string `json:"api_key"`
	Name      string `json:"name"`
	ExpiresAt string `json:"expires_at,omitempty"`
	Scopes    string `json:"scopes,omitempty"`
}

type UpdateAPIKeyRequest struct {
	ExpiresAt *string `json:"expires_at,omitempty"`
	Scopes    *string `json:"scopes,omitempty"`
}

type APIKeyResponse struct {
	APIKey    string `json:"api_key"`
	Name      string `json:"name"`
	ExpiresAt string `json:"expires_at,omitempty"`
	Scopes    string `json:"scopes,omitempty"`
}
