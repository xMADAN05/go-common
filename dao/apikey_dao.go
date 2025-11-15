package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/xMADAN05/go-common/common/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type APIKeyDAO struct {
	client    *dynamodb.Client
	tableName string
}

func NewDAO(tableName string, region string) (*APIKeyDAO, error) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return &APIKeyDAO{
		client:    client,
		tableName: tableName,
	}, nil
}

func (dao *APIKeyDAO) Create(ctx context.Context, record models.APIKeyRecord) error {
	item, err := attributevalue.MarshalMap(record)
	if err != nil {
		return fmt.Errorf("failed to marshall api key record: %w", err)
	}

	input := &dynamodb.PutItemInput{

		TableName: aws.String(dao.tableName),
		Item:      item,
	}

	_, err = dao.client.PutItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to put item: %w", err)
	}

	return nil
}

func (dao *APIKeyDAO) GetByID(ctx context.Context, apiKey string) (*models.APIKeyRecord, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(dao.tableName),
		Key: map[string]types.AttributeValue{
			"api_key": &types.AttributeValueMemberS{Value: apiKey},
		},
	}

	res, err := dao.client.GetItem(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	if res.Item == nil {
		return nil, fmt.Errorf("api key not found: %w", err)
	}

	var record models.APIKeyRecord

	if err := attributevalue.UnmarshalMap(res.Item, &record); err != nil {
		return nil, fmt.Errorf("failed to unmarshal recrod: %w", err)
	}

	return &record, nil

}
