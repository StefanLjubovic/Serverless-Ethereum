package repository

import (
	"backend/dto"
	"backend/model"
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
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

func (userRepository *UsersDynamoDBStore) GetByUsername(username string) (*model.User, error) {
	err := userRepository.GetDBClient()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(userRepository.TableName),
		Key: map[string]types.AttributeValue{
			"Username": &types.AttributeValueMemberS{Value: username},
		},
	}

	result, err := userRepository.DynamoDbClient.GetItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Failed to get course: ", err)
		return nil, err
	}

	user := new(model.User)
	err = attributevalue.UnmarshalMap(result.Item, user)
	if err != nil {
		fmt.Println("Failed to unmarshal course")
		return nil, err
	}

	return user, nil
}

func (userRepository *UsersDynamoDBStore) UpdateUsersCourses(id string, usersCourses *[]model.UsersCourse) error {
	courseList, err := attributevalue.MarshalList(usersCourses)
	if err != nil {
		return err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(userRepository.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET users_courses = :users_courses"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":users_courses": &types.AttributeValueMemberL{Value: courseList},
		},
	}

	_, err = userRepository.DynamoDbClient.UpdateItem(context.TODO(), updateInput)

	if err != nil {
		fmt.Println("Error updating course item: ", err)
		return err
	}

	return nil
}

func (userRepository *UsersDynamoDBStore) GetUserCourses(username string) (*[]model.UsersCourse, error) {
	err := userRepository.GetDBClient()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(userRepository.TableName),
		Key: map[string]types.AttributeValue{
			"Username": &types.AttributeValueMemberS{Value: username},
		},
	}

	result, err := userRepository.DynamoDbClient.GetItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Failed to get course: ", err)
		return nil, err
	}

	user := new(model.User)
	err = attributevalue.UnmarshalMap(result.Item, user)
	if err != nil {
		fmt.Println("Failed to unmarshal course")
		return nil, err
	}

	return &user.UsersCourses, nil
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
