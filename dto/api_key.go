package dto

type CreateAPIKeyRequest struct {
	WorkspaceID   string `dynamodbav:"workspace_id"`
	ProjectID     string `dynamodbav:"project_id"`
	ApplicationID string `dynamodbav:"application_id"`
	EnvironmentID string `dynamodbav:"environment_id"`

	Scopes    string `dynamodbav:"scopes"`
	Active    string `dynamodbav:"active"`
	ExpiresAt string `dynamodbav:"expires_at"`
}

type CreateAPIKeyResponse struct {
	APIKey   string `json:"api_key"`
	APIKeyID string `dynamodbav:"api_key_id"`

	WorkspaceID   string `dynamodbav:"workspace_id"`
	ProjectID     string `dynamodbav:"project_id"`
	ApplicationID string `dynamodbav:"application_id"`
	EnvironmentID string `dynamodbav:"environment_id"`

	Scopes    string `dynamodbav:"scopes"`
	Active    string `dynamodbav:"active"`
	ExpiresAt string `dynamodbav:"expires_at"`
	CreatedAt string `dynamodbav:"created_at"`
}

// Use *string to distinguish between “not provided” (nil) and “provided” (even if empty),
// enabling proper PATCH behavior without overwriting fields unintentionally.
type UpdateAPIKeyRequest struct {
	Scopes    *string `json:"scopes,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
}

type APIKeyResponse struct {
	APIKeyHash string `json:"api_key_hash"`
	APIKeyID   string `dynamodbav:"api_key_id"`

	WorkspaceID   string `dynamodbav:"workspace_id"`
	ProjectID     string `dynamodbav:"project_id"`
	ApplicationID string `dynamodbav:"application_id"`
	EnvironmentID string `dynamodbav:"environment_id"`

	Scopes    string `dynamodbav:"scopes"`
	Active    string `dynamodbav:"active"`
	ExpiresAt string `dynamodbav:"expires_at"`
	CreatedAt string `dynamodbav:"created_at"`
}
