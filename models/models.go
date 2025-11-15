package models

type APIKeyRecord struct {
	APIKey      string `dynamodbav:"api_key"`
	ServiceName string `dynamodbav:"service_name"`
	Active      string `dynamodbav:"active"`
	Scopes      string `dynamodbav:"scopes"`
	ExpiresAt   string `dynamodbav:"expires_at"`
}
