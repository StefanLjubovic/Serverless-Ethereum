package lambdas

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	services "../../service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var courseService services.CoursesService

func init() {
	fmt.Printf("Init")
	serviceString := os.Getenv("USER_SERVICE")
	if serviceString == "" {
		log.Fatal("missing environment variable USER_SERVICE")
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

	courses, err := courseService.GetAllCourses()

	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, nil
	}
	responseBody, err := json.Marshal(courses)
	if err != nil {
		log.Printf("failed to marshal courses: %v", err)
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
