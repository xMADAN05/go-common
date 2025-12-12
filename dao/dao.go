package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoRepository[T any] struct {
	client    *dynamodb.Client
	tableName string
}

func NewRepository[T any](tableName string, region string) (*DynamoRepository[T], error) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoRepository[T]{
		client:    client,
		tableName: tableName,
	}, nil
}

func NewRepositoryWithClient[T any](client *dynamodb.Client, tableName string) *DynamoRepository[T] {
	return &DynamoRepository[T]{client, tableName}
}

func (r *DynamoRepository[T]) Put(ctx context.Context, item T) error {
	val, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("failed to marshall item: %w", err)
	}

	input := &dynamodb.PutItemInput{

		TableName: aws.String(r.tableName),
		Item:      val,
	}

	_, err = r.client.PutItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to put item: %w", err)
	}

	return nil
}

func (r *DynamoRepository[T]) Get(ctx context.Context, key map[string]types.AttributeValue) (*T, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key:       key,
	}

	res, err := r.client.GetItem(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	if res.Item == nil {
		return nil, fmt.Errorf("api key not found: %w", err)
	}

	var record T

	if err := attributevalue.UnmarshalMap(res.Item, &record); err != nil {
		return nil, fmt.Errorf("failed to unmarshal recrod: %w", err)
	}

	return &record, nil

}

func (r *DynamoRepository[T]) GetAllRecords(ctx context.Context) ([]T, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}
	var res []T
	paginator := dynamodb.NewScanPaginator(r.client, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to scan table :%w", err)
		}
		var items []T
		if err := attributevalue.UnmarshalListOfMaps(page.Items, &items); err != nil {
			return nil, fmt.Errorf("failed to unmarshal items :%w", err)
		}
		res = append(res, items...)
	}
	return res, nil
}

func (r *DynamoRepository[T]) Update(
	ctx context.Context,
	key map[string]types.AttributeValue,
	updateExpr string,
	exprAttrValues map[string]types.AttributeValue,
	exprAttrNames map[string]string,
) error {
	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(r.tableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpr),
		ExpressionAttributeValues: exprAttrValues,
	}
	if exprAttrNames != nil {
		input.ExpressionAttributeNames = exprAttrNames
	}
	_, err := r.client.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}
	return nil
}

func (r *DynamoRepository[T]) Delete(ctx context.Context, key map[string]types.AttributeValue) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key:       key,
	}
	_, err := r.client.DeleteItem(ctx, input)

	if err != nil {
		log.Fatalf("failed to delete item: %v", err)
	}

	return err
}
