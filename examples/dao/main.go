package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/xMADAN05/go-common/dao"
	"github.com/xMADAN05/go-common/models"
	"github.com/xMADAN05/go-common/utils"
)

func main() {
	// dynamodbutils.CreateTable("api_keys_table")
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-2"))
	if err != nil {
		log.Fatalf("unable to load AWS Config : %v", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	apiKeyDAO := dao.NewAPIKeyDAO(client, "api_keys_table")

	apiKey := "supersecretkey"
	hashed := utils.HashKey(apiKey)
	apiKeyID := uuid.NewString()
	record := models.APIKeyRecord{

		PK:            "APIKEY#" + apiKeyID,
		SK:            "METADATA#" + apiKeyID,
		APIKeyHash:    hashed,
		WorkspaceID:   "ws_01",
		ProjectID:     "proj_01",
		EnvironmentID: "dev",
		ApplicationID: "chat_app",
		APIKeyID:      apiKeyID,
		Active:        "true",
		Scopes:        "read, write",
		CreatedAt:     time.Now().Format(time.RFC3339),
		ExpiresAt:     "2026-12-31T23:59:59Z",
	}

	if err := apiKeyDAO.Put(ctx, record); err != nil {
		log.Fatalf("failed to put API key: %v", err)
	}

	rec, err := apiKeyDAO.Get(ctx, apiKeyID)
	if err != nil {
		log.Fatalf("failed to get API key: %v", err)
	} else {
		fmt.Printf("record : %v\n", rec)
	}

	// if err := apiKeyDAO.Delete(ctx, "supersecretkey"); err != nil {
	// 	log.Fatalf("failed to delete API key: %v", err)
	// } else {
	// 	fmt.Printf("API key deleted successfully!")
	// }
}
