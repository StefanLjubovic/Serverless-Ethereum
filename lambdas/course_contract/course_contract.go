package main

import (
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
	fmt.Println("started")
	param := req.PathParameters["price"]
	floatNumber, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusBadRequest}, nil

	}
	payloadDTO, err := courseService.DeployContract(floatNumber)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, nil

	}
	ret, _ := json.Marshal(payloadDTO)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       string(ret),
	}, nil
}

func main() {
	lambda.Start(handler)
}
