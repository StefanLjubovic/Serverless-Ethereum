package main

import (
	"backend/dto"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func init() {

}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var body dto.SignInDto
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

	signInInput := cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String("3s9evb0dc0qspc503fnbnajgqm"),
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(body.Username),
			"PASSWORD": aws.String(body.Password),
		},
	}

	signInResponse, err := cognitoClient.InitiateAuth(&signInInput)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Error signing in: %s", err),
		}, nil
	}

	responseBody, _ := json.Marshal(dto.SignInResponseDto{
		Message: "User signed in successfully",
		Token:   aws.StringValue(signInResponse.AuthenticationResult.IdToken),
	})

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
	}, nil
}
