package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/xMADAN05/go-common/dao"
	"github.com/xMADAN05/go-common/models"
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

	record := models.APIKeyRecord{
		APIKey:      "supersecretkey",
		ServiceName: "Test",
		Active:      "true",
		Scopes:      "read, write",
		ExpiresAt:   "2026-12-31T23:59:59Z",
	}

	if err := apiKeyDAO.Put(ctx, record); err != nil {
		log.Fatalf("failed to put API key: %v", err)
	}

	rec, err := apiKeyDAO.Get(ctx, "supersecretkey")
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
