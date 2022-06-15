package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (db *DB) newDynamoDB() error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	db.DDB = dynamodb.New(sess)
	return nil
}

func (db *DB) putItem(url string) error {
	randID := strconv.Itoa(rand.Intn(100))
	epoch := strconv.FormatInt(time.Now().Unix(), 10)
	input := &dynamodb.PutItemInput{
		TableName: aws.String("table"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(randID),
			},
			"time": {
				N: aws.String(epoch),
			},
			"url": {
				S: aws.String(url),
			},
		},
	}
	if _, err := db.DDB.PutItem(input); err != nil {
		return err
	}
	return nil
}
