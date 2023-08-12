package service

import (
	"backend/dto"
	model "backend/model"
	"backend/util"
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type CoursesService struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
	S3ImageService S3ImageServiceImpl
}

func NewCoursesHandler(bucketName string) *CoursesService {
	s3Service := NewS3ClientService(bucketName)
	return &CoursesService{
		TableName:      "Course",
		DynamoDbClient: nil,
		S3ImageService: *s3Service,
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

func (coursesService *CoursesService) GetCourseById(id string) (*model.Course, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		fmt.Println("Failed to make configuration: ", err)
	}
	coursesService.DynamoDbClient = dynamodb.NewFromConfig(cfg)

	input := &dynamodb.GetItemInput{
		TableName: aws.String(coursesService.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: id},
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

func (coursesService *CoursesService) UploadObject(image []byte) (string, error) {
	coursesService.S3ImageService.Start(context.TODO())
	return coursesService.S3ImageService.UploadObject(context.TODO(), image)
}

func (coursesService *CoursesService) Save(dto dto.CourseCreate) error {
	ethPrice, err := util.ConvertUSDToETH(dto.PriceUSD)
	fmt.Println(ethPrice)
	if err != nil {
		fmt.Println("Failed to convert usd to eth: ", err)
		return err
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		fmt.Println("Failed to make configuration: ", err)
		return err
	}
	coursesService.DynamoDbClient = dynamodb.NewFromConfig(cfg)
	item := map[string]types.AttributeValue{
		"id":          &types.AttributeValueMemberN{Value: strconv.FormatUint(uint64(dto.ID), 10)},
		"name":        &types.AttributeValueMemberS{Value: dto.Name},
		"description": &types.AttributeValueMemberS{Value: dto.Description},
		"certificate": &types.AttributeValueMemberS{Value: dto.Certificate},
		"image":       &types.AttributeValueMemberS{Value: dto.Image},
		"price": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
			"priceETH": &types.AttributeValueMemberN{Value: strconv.FormatFloat(dto.PriceUSD, 'f', -1, 64)},
			"priceUSD": &types.AttributeValueMemberN{Value: strconv.FormatFloat(ethPrice, 'f', -1, 64)},
		}},
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(coursesService.TableName),
		Item:      item,
	}
	_, err = coursesService.DynamoDbClient.PutItem(context.TODO(), input)
	if err != nil {
		fmt.Println("error inserting to database: ", err)
		return err
	}
	return nil
}

func (coursesService *CoursesService) DeployContract(price float64) (*dto.CourseContractResp, error) {
	id := util.GenerateRowID(1)
	priceInWei, err := util.ConvertUSDToWei(price)
	if err != nil {
		return nil, err
	}
	retVal := dto.CourseContractResp{
		PriceInWei: priceInWei,
		ID:         id,
	}
	return &retVal, nil
}
