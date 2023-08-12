package main

import (
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var courseService service.CoursesService

func init() {
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

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	fileContent := []byte(req.Body)

	filename, err := courseService.UploadObject(fileContent)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: 500}, err
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(filename),
	}, nil
}
