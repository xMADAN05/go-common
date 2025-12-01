package models

// type APIKeyRecord struct {
// 	APIKeyID    string `dynamodbav:"api_key_id`
// 	APIKeyHash  string `dynamodbav:"api_key_hash"`
// 	ServiceName string `dynamodbav:"service_name"`
// 	Scopes      string `dynamodbav:"scopes"`
// 	Active      string `dynamodbav:"active"`
// 	ExpiresAt   string `dynamodbav:"expires_at"`
// 	CreatedAt   string `dynamodbav:"created_at"`
// }

type APIKeyRecord struct {
	PK         string `dynamodbav:"PK"`
	SK         string `dynamodbav:"SK"`
	APIKeyID   string `dynamodbav:"api_key_id"`
	APIKeyHash string `dynamodbav:"api_key_hash"`

	WorkspaceID   string `dynamodbav:"workspace_id"`
	ProjectID     string `dynamodbav:"project_id"`
	ApplicationID string `dynamodbav:"application_id"`
	EnvironmentID string `dynamodbav:"environment_id"`

	Scopes    string `dynamodbav:"scopes"`
	Active    string `dynamodbav:"active"`
	ExpiresAt string `dynamodbav:"expires_at"`
	CreatedAt string `dynamodbav:"created_at"`
}
