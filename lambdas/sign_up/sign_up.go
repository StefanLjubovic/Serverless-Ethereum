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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var userService service.UsersService

func init() {
	serviceString := os.Getenv("USERS_SERVICE")
	if serviceString == "" {
		log.Fatal("Missing environment variable USERS_SERVICE")
		return
	}

	err := json.Unmarshal([]byte(serviceString), &userService)
	if err != nil {
		fmt.Println("Error unmarshaling user service")
		return
	}
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var body dto.UserCreate
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	signUpInput := cognitoidentityprovider.SignUpInput{
		ClientId: aws.String("3s9evb0dc0qspc503fnbnajgqm"),
		Username: &body.Username,
		Password: &body.Password,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(body.Email),
			},
		},
	}

	fmt.Println(body)

	_, err = cognitoClient.SignUp(&signUpInput)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Error signing up: %s", err),
		}, nil
	}

	fmt.Println(body)
	err = userService.Save(body)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusCreated}, nil
}
