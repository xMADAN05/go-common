package dynamodbutils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateTable(tableName string) error {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-2"))
	if err != nil {
		return fmt.Errorf("unable to load SDK config, %w", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("api_key"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("api_key"),
				KeyType:       types.KeyTypeHash,
			},
		},

		BillingMode: types.BillingModePayPerRequest,
	}

	_, err = client.CreateTable(ctx, input)
	if err != nil {
		var existsErr *types.ResourceInUseException
		if errors.As(err, &existsErr) {
			fmt.Printf("Table `%s` already exists\n", tableName)
		} else {
			return fmt.Errorf("failed to create table: %w", err)
		}
	} else {
		fmt.Printf("Creating table `%s`.\n", tableName)
		waiter := dynamodb.NewTableExistsWaiter(client)
		err = waiter.Wait(ctx, &dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		}, 5*time.Minute)

		if err != nil {
			return fmt.Errorf("error waiting for table creation: %w", err)
		}

		resp, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		})

		if err != nil {
			return fmt.Errorf("describe table failed: %w", err)
		}

		fmt.Printf("Table status: %v\n", resp.Table.TableStatus)
	}

	return nil

}
