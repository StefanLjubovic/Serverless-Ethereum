package util

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func ExtractUsername(req events.APIGatewayV2HTTPRequest) (string, error) {

	authHeader := req.Headers["authorization"]
	token := strings.TrimPrefix(authHeader, "Bearer ")
	payloadParts := strings.Split(token, ".")
	if len(payloadParts) != 3 {
		return "", nil
	}
	payload, err := base64.RawStdEncoding.DecodeString(payloadParts[1])
	if err != nil {
		return "", err
	}
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return "", err
	}
	username, ok := claims["cognito:username"].(string)
	if !ok {
		return "", err
	}
	return username, nil
}
