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

	// sess, err := session.NewSession()
	// if err != nil {
	// 	return err
	// }
	// db := dynamodb.New(sess)

	return nil

}

func main() {
	lambda.Start(HandleRequest)
}
