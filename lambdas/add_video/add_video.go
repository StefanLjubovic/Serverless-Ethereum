package main

import (
	"backend/dto"
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

var courseService service.CoursesService

func init() {
	fmt.Println("Init")
	serviceString := os.Getenv("COURSES_SERVICE")
	if serviceString == "" {
		log.Fatal("Missing environment variable COURSES_SERVICE")
		return
	}
	err := json.Unmarshal([]byte(serviceString), &courseService)
	if err != nil {
		fmt.Println("Error unmarshaling course service")
		return
	}
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var body dto.AddVideoDTO
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusBadRequest}, nil

	}

	err = courseService.AddVideo(body)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusNotFound}, nil

	}
	return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(handler)
}
