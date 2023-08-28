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

type CoursesDynamoDBStore struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewCoursesDBStore(tableName string) *CoursesDynamoDBStore {
	return &CoursesDynamoDBStore{
		DynamoDbClient: nil,
		TableName:      tableName,
	}
}

func (coursesRepository *CoursesDynamoDBStore) GetAllCourses() (*[]model.Course, error) {
	err := coursesRepository.GetDBClient()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	input := &dynamodb.ScanInput{
		TableName: aws.String(coursesRepository.TableName),
	}

	result, err := coursesRepository.DynamoDbClient.Scan(context.TODO(), input)
	if err != nil {
		fmt.Println("Failed to scan all courses:", err)
		return nil, err
	}
	var courses []model.Course
	for _, item := range result.Items {
		var course model.Course
		err = attributevalue.UnmarshalMap(item, &course)
		if err != nil {
			continue
		}
		courses = append(courses, course)
	}
	return &courses, nil
}

func (coursesRepository *CoursesDynamoDBStore) GetCourseById(id string) (*model.Course, error) {
	err := coursesRepository.GetDBClient()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(coursesRepository.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: id},
		},
	}

	result, err := coursesRepository.DynamoDbClient.GetItem(context.TODO(), input)
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

func (coursesRepository *CoursesDynamoDBStore) Save(ethPrice float64, dto dto.CourseCreate) error {
	err := coursesRepository.GetDBClient()
	if err != nil {
		return err
	}
	item := map[string]types.AttributeValue{
		"id":          &types.AttributeValueMemberN{Value: strconv.FormatUint(uint64(dto.ID), 10)},
		"name":        &types.AttributeValueMemberS{Value: dto.Name},
		"description": &types.AttributeValueMemberS{Value: dto.Description},
		"certificate": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
			"name":        &types.AttributeValueMemberS{Value: dto.Certificate.Name},
			"description": &types.AttributeValueMemberS{Value: dto.Certificate.Description},
			"image_path":  &types.AttributeValueMemberS{Value: dto.Certificate.ImagePath},
		}},
		"image": &types.AttributeValueMemberS{Value: dto.Image},
		"price": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
			"price_usd": &types.AttributeValueMemberN{Value: strconv.FormatFloat(dto.PriceUSD, 'f', -1, 64)},
			"price_eth": &types.AttributeValueMemberN{Value: strconv.FormatFloat(ethPrice, 'f', -1, 64)},
		}},
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(coursesRepository.TableName),
		Item:      item,
	}
	_, err = coursesRepository.DynamoDbClient.PutItem(context.TODO(), input)
	if err != nil {
		fmt.Println("error inserting to database: ", err)
		return err
	}
	return nil
}

func (coursesRepository *CoursesDynamoDBStore) UpdateSections(sections []model.Section, courseId int) error {

	err := coursesRepository.GetDBClient()
	if err != nil {
		return err
	}
	courseList, err := attributevalue.MarshalList(sections)
	if err != nil {
		return err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(coursesRepository.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.FormatUint(uint64(courseId), 10)},
		},
		UpdateExpression: aws.String("SET sections = :sections"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":sections": &types.AttributeValueMemberL{Value: courseList},
		},
	}

	_, err = coursesRepository.DynamoDbClient.UpdateItem(context.TODO(), updateInput)

	if err != nil {
		fmt.Println("Error updating course item: ", err)
		return err
	}

	return nil
}

func (coursesRepository *CoursesDynamoDBStore) GetDBClient() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		fmt.Println("Failed to make configuration: ", err)
		return err
	}
	coursesRepository.DynamoDbClient = dynamodb.NewFromConfig(cfg)
	return nil
}
