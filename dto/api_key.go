package dto

type CreateAPIKeyRequest struct {
	ServiceName string `json:"service_name" binding:"required"`
	ExpiresAt   string `json:"expires_at,omitempty"`
	Scopes      string `json:"scopes,omitempty"`
}

type CreateAPIKeyResponse struct {
	APIKey      string `json:"api_key"`
	ServiceName string `json:"service_name"`
	ExpiresAt   string `json:"expires_at,omitempty"`
	Scopes      string `json:"scopes,omitempty"`
}

type UpdateAPIKeyRequest struct {
	ExpiresAt *string `json:"expires_at,omitempty"`
	Scopes    *string `json:"scopes,omitempty"`
}

type APIKeyResponse struct {
	APIKey      string `json:"api_key"`
	ServiceName string `json:"service_name"`
	ExpiresAt   string `json:"expires_at,omitempty"`
	Scopes      string `json:"scopes,omitempty"`
}
