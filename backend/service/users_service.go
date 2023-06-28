package service

import (
	model "backend/model"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/jsii-runtime-go"
)

type UsersService struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewUserHandler() *UsersService {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return &UsersService{
		TableName:      "User",
		DynamoDbClient: dynamodb.NewFromConfig(sdkConfig),
	}

}

func (usersService *UsersService) GetUserByUsername(username string) (*model.User, error) {

	response, err := usersService.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"Username": &types.AttributeValueMemberS{
				Value: username,
			},
		},

		TableName: jsii.String(usersService.TableName),
	})
	if err != nil {
		return nil, err
	}
	user := model.User{}
	err = attributevalue.UnmarshalMap(response.Item, &user)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return nil, nil
}
