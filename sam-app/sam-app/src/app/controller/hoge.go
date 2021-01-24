package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var db *dynamodb.DynamoDB
var dbEndpoint = "http://host.docker.internal:4569"
var region = "ap-northeast-1"
var testTable = "hoge_test"

type CompanyResponse struct {
	Company string `json:"company"`
	Year    string `json:"year"`
}

func write(ctx context.Context, tableName string, v interface{}) error {
	av, err := dynamodbattribute.MarshalMap(v)

	if err != nil {
		return fmt.Errorf("dynamodb attribute marshalling map: %w", err)
	}
	i := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	if _, err = db.PutItemWithContext(ctx, i); err != nil {
		return fmt.Errorf("dynamodb put item: %w", err)
	}
	return nil
}

type HogeController struct {}

func (hc HogeController) Index(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(dbEndpoint),
		Region:   aws.String(region),
	}))
	db = dynamodb.New(sess)

	response := CompanyResponse{
		Company: "Future",
		Year:    "1989",
	}

	jsonBytes, _ := json.Marshal(response)

	fmt.Println("--------------ok1")

	if err := write(ctx, testTable, response); err != nil {
		fmt.Print("%s", err)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}