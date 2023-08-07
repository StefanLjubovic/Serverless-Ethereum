package service

import (
	model "backend/model"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type CoursesService struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewCoursesHandler() *CoursesService {
	//sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	//if err != nil {
	//	log.Fatalf("unable to load SDK config, %v", err)
	//}
	return &CoursesService{
		TableName:      "Course",
		DynamoDbClient: nil,
	}

}

func (coursesService *CoursesService) GetAllCourses() (*[]model.Course, error) {

	input := &dynamodb.QueryInput{
		TableName:              aws.String(coursesService.TableName),
		KeyConditionExpression: aws.String("id = :id_value"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id_value": &types.AttributeValueMemberN{Value: "1"},
		},
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		fmt.Println("Error with cfg: ", err)
	}
	coursesService.DynamoDbClient = dynamodb.NewFromConfig(cfg)

	result, err := coursesService.DynamoDbClient.Query(context.TODO(), input)
	if err != nil {
		fmt.Println("Failed to query DynamoDB:", err)
		return nil, err
	}
	var courses []model.Course
	for _, item := range result.Items {
		var course model.Course
		// Access individual attributes of each item
		err = attributevalue.UnmarshalMap(item, &course)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return &courses, nil
}
