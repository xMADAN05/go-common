package dao

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/xMADAN05/go-common/models"
)

type APIKeyDAO struct {
	repo *DynamoRepository[models.APIKeyRecord]
}

func NewAPIKeyDAO(client *dynamodb.Client, tableName string) *APIKeyDAO {

	return &APIKeyDAO{
		repo: &DynamoRepository[models.APIKeyRecord]{client, tableName},
	}
}

func (d *APIKeyDAO) Put(ctx context.Context, Item models.APIKeyRecord) error {
	return d.repo.Put(ctx, Item)
}

func (d *APIKeyDAO) Get(ctx context.Context, apiKey string) (*models.APIKeyRecord, error) {
	key := map[string]types.AttributeValue{
		"api_key": &types.AttributeValueMemberS{Value: apiKey},
	}
	return d.repo.Get(ctx, key)
}

func (d *APIKeyDAO) Update(
	ctx context.Context,
	key map[string]types.AttributeValue,
	updateExpr string,
	exprAttrValues map[string]types.AttributeValue,
	exprAttrNames map[string]string,
) error {
	return d.repo.Update(ctx, key, updateExpr, exprAttrValues, exprAttrNames)
}

func (d *APIKeyDAO) Delete(ctx context.Context, apiKey string) error {
	key := map[string]types.AttributeValue{
		"api_key": &types.AttributeValueMemberS{Value: apiKey},
	}
	return d.repo.Delete(ctx, key)
}
