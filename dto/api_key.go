package dto

type CreateAPIKeyRequest struct {
	ServiceName string `json:"service_name" binding:"required"`
	Scopes      string `json:"scopes,omitempty"`
	ExpiresAt   string `json:"expires_at,omitempty"`
}

type CreateAPIKeyResponse struct {
	APIKey      string `json:"api_key"`
	ServiceName string `json:"service_name"`
	Scopes      string `json:"scopes,omitempty"`
	Active      string `json:"active"`
	CreatedAt   string `json:"created_at"`
	ExpiresAt   string `json:"expires_at,omitempty"`
}

// Use *string to distinguish between “not provided” (nil) and “provided” (even if empty),
// enabling proper PATCH behavior without overwriting fields unintentionally.
type UpdateAPIKeyRequest struct {
	Scopes    *string `json:"scopes,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
}

type APIKeyResponse struct {
	APIKey      string `json:"api_key"`
	ServiceName string `json:"service_name"`
	Scopes      string `json:"scopes,omitempty"`
	Active      string `json:"active"`
	CreatedAt   string `json:"created_at"`
	ExpiresAt   string `json:"expires_at,omitempty"`
}
