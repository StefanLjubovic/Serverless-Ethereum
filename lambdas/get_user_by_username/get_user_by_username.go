package main

import (
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var usersService service.UsersService

func init() {
	fmt.Println("Init")
	serviceString := os.Getenv("USERS_SERVICE")
	if serviceString == "" {
		log.Fatal("Missing environment variable USERS_SERVICE")
		return
	}
	err := json.Unmarshal([]byte(serviceString), &usersService)
	if err != nil {
		fmt.Println("Error unmarshaling course service")
		return
	}
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	username := "ljubo"
	user, err := usersService.GetUserByUsername(username)
	fmt.Println(user)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusNotFound,
			Body:       string(err.Error()),
		}, nil
	}

	responseBody, err := json.Marshal(user)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
