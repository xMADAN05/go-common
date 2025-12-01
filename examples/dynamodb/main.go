package main

import (
	dynamodbutils "github.com/xMADAN05/go-common/utils/dynamodb"
)

func main() {

	dynamodbutils.CreateTable("api_keys_table")
}
