package main

import (
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"lambdas/util"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var usersService service.UsersService

func init() {
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

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	username, err := util.ExtractUsername(req)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusUnauthorized}, nil
	}
	courses, err := usersService.GetUsersCourses(username)
	fmt.Println(courses)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusBadRequest}, nil
	}

	responseBody, err := json.Marshal(courses)
	fmt.Println(responseBody)
	if err != nil {
		log.Printf("Failed to marshal courses: %v", err)
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
