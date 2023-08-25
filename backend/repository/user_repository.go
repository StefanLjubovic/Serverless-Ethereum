package repository

import (
	"backend/dto"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"strconv"
)

type UsersDynamoDBStore struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewUsersDBStore(tableName string) *UsersDynamoDBStore {
	return &UsersDynamoDBStore{
		DynamoDbClient: nil,
		TableName:      tableName,
	}
}

func (userRepository *UsersDynamoDBStore) Save(dto dto.UserCreate) error {
	err := userRepository.GetDBClient()
	if err != nil {
		return err
	}
	item := map[string]types.AttributeValue{
		"id":       &types.AttributeValueMemberN{Value: strconv.FormatUint(uint64(dto.ID), 10)},
		"name":     &types.AttributeValueMemberS{Value: dto.Name},
		"surname":  &types.AttributeValueMemberS{Value: dto.Surname},
		"username": &types.AttributeValueMemberS{Value: dto.Username},
		"email":    &types.AttributeValueMemberS{Value: dto.Email},
		"image":    &types.AttributeValueMemberS{Value: ""},
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(userRepository.TableName),
		Item:      item,
	}
	_, err = userRepository.DynamoDbClient.PutItem(context.TODO(), input)
	if err != nil {
		fmt.Println("error inserting to database: ", err)
		return err
	}

	return nil
}

func (userRepository *UsersDynamoDBStore) GetDBClient() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		fmt.Println("Failed to make configuration: ", err)
		return err
	}
	userRepository.DynamoDbClient = dynamodb.NewFromConfig(cfg)
	return nil
}
