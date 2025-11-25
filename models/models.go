package models

type APIKeyRecord struct {
	APIKey      string `dynamodbav:"api_key"`
	ServiceName string `dynamodbav:"service_name"`
	Scopes      string `dynamodbav:"scopes"`
	Active      string `dynamodbav:"active"`
	ExpiresAt   string `dynamodbav:"expires_at"`
	CreatedAt   string `dynamodbav:"created_at"`
}
