package main

import (
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/grokify/go-awslambda"
)

var courseService service.CoursesService

func init() {
	serviceString := os.Getenv("COURSES_SERVICE")
	if serviceString == "" {
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

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayV2HTTPResponse, error) {
	r, err := awslambda.NewReaderMultipart(req)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, err
	}
	part, err := r.NextPart()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, err
	}
	content, err := io.ReadAll(part)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, err
	}
	type Custom struct {
		FileName string `json:"filename"`
		File     []byte `json:"file"`
	}

	custom := Custom{
		File:     content,
		FileName: part.FileName(),
	}
	path, err := courseService.UploadObject(custom.File, custom.FileName)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       path,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
