package main

import (
	"backend/dto"
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
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
		return
	}
	err := json.Unmarshal([]byte(serviceString), &usersService)
	if err != nil {
		return
	}
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	username := "ljubo"
	var body dto.AddWatchedDTO
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusBadRequest}, nil

	}
	err = usersService.AddWatchedVideo(username, &body)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusNotFound,
			Body:       string("Course with provided id does not exist"),
		}, nil
	}
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
