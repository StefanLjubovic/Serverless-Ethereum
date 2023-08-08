package main

import (
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"os"
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
	param := req.PathParameters["id"]
	course, err := courseService.GetCourseById(param)
	fmt.Println(course)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusNotFound,
			Body:       string("Course with provided id does not exist"),
		}, nil
	}

	responseBody, err := json.Marshal(course)
	fmt.Println(responseBody)
	if err != nil {
		log.Printf("Failed to marshal course: %v", err)
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
