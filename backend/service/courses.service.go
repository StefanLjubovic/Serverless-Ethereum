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
	"strconv"
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
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		fmt.Println("Failed to make configuration: ", err)
	}
	coursesService.DynamoDbClient = dynamodb.NewFromConfig(cfg)

	input := &dynamodb.ScanInput{
		TableName: aws.String(coursesService.TableName),
	}

	result, err := coursesService.DynamoDbClient.Scan(context.TODO(), input)
	if err != nil {
		fmt.Println("Failed to scan all courses:", err)
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

func (coursesService *CoursesService) GetCourseById(id int) (*model.Course, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		fmt.Println("Failed to make configuration: ", err)
	}
	coursesService.DynamoDbClient = dynamodb.NewFromConfig(cfg)

	input := &dynamodb.GetItemInput{
		TableName: aws.String(coursesService.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
		},
	}

	result, err := coursesService.DynamoDbClient.GetItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Failed to get course: ", err)
		return nil, err
	}

	course := new(model.Course)
	err = attributevalue.UnmarshalMap(result.Item, course)
	if err != nil {
		fmt.Println("Failed to unmarshal course")
		return nil, err
	}

	return course, nil
}
