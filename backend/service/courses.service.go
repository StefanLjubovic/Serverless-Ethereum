package service

import (
	model "backend/model"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type CoursesService struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewCoursesHandler() *CoursesService {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return &CoursesService{
		TableName:      "Course",
		DynamoDbClient: dynamodb.NewFromConfig(sdkConfig),
	}

}

func (coursesService *CoursesService) GetAllCourses() (*[]model.Course, error) {

	input := &dynamodb.QueryInput{
		TableName: jsii.String(coursesService.TableName),
	}
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
