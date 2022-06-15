package main

import "github.com/aws/aws-sdk-go/service/dynamodb"

type Payload struct {
	Body string `json:"body"`
}

type Request struct {
	URL string `json:"url"`
}

type DBAPI interface {
	putItem(url string) error
}

type DB struct {
	DDB   *dynamodb.DynamoDB
	DBAPI DBAPI
}
