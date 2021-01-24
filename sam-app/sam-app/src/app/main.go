package main

import (
	"app/controller"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch {
		case request.HTTPMethod == "GET" && request.Path == "/hello":
			return controller.HogeController{}.Index(request)
	}

	jsonBytes, _ := json.Marshal(map[string]interface{}{"error" : "path error"})
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 400,
	}, nil
}

func main() {
	fmt.Println(os.Getenv("PARAM1"))
	lambda.Start(handler)
}
