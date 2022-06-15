package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, payload Payload) error {
	url, err := parseRequest(payload)
	if err != nil {
		return err
	}
	isAmazonURL, err := checkAmazonURL(url)
	if err != nil {
		return err
	}
	if !isAmazonURL {
		return fmt.Errorf("this is not amazon url: %v: %v", err, url)
	}

	db := new(DB)
	if err := db.newDynamoDB(); err != nil {
		return err
	}
	if err := db.putItem(url); err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
