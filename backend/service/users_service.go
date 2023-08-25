package service

import (
	"backend/dto"
	model "backend/model"
	"backend/repository"
	"context"
	"fmt"
	"github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UsersService struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
	UserRepository repository.UsersDynamoDBStore
}

func NewUserHandler() *UsersService {
	//sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	//if err != nil {
	//	log.Fatalf("unable to load SDK config, %v", err)
	//}
	repository := repository.NewUsersDBStore("User")
	return &UsersService{
		TableName:      "User",
		DynamoDbClient: nil,
		UserRepository: *repository,
	}

}

func (usersService *UsersService) GetUserByUsername(username string) (*model.User, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		fmt.Println("Failed to make configuration: ", err)
	}
	usersService.DynamoDbClient = dynamodb.NewFromConfig(cfg)

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
	return &user, nil
}

func (usersService *UsersService) Save(dto dto.UserCreate) error {
	return usersService.UserRepository.Save(dto)
}

// func (usersService *UsersService) AddUserCourse(username string, id uint64) error {
// 	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
// 	if err != nil {
// 		fmt.Println("Failed to make configuration: ", err)
// 		return err
// 	}
// 	user, err := usersService.GetUserByUsername(username)
// 	user.CreatedCourses = append(user.CreatedCourses, uint(id))
// 	if err != nil {
// 		fmt.Println("Failed to find user: ", err)
// 		return err
// 	}

// 	usersService.DynamoDbClient = dynamodb.NewFromConfig(cfg)

// 	courseList := make([]types.AttributeValue, len(user.CreatedCourses))
// 	for _, id := range user.CreatedCourses {
// 		idStr := strconv.FormatUint(uint64(id), 10)
// 		courseList = append(courseList, &types.AttributeValueMemberS{Value: idStr})
// 	}

// 	updateInput := &dynamodb.UpdateItemInput{
// 		TableName: aws.String(usersService.TableName),
// 		Key: map[string]types.AttributeValue{
// 			"username": &types.AttributeValueMemberS{Value: username},
// 		},
// 		UpdateExpression: aws.String("SET created_courses = :courses"),
// 		ExpressionAttributeValues: map[string]types.AttributeValue{
// 			":courses": &types.AttributeValueMemberL{Value: courseList},
// 		},
// 	}

// 	_, err = usersService.DynamoDbClient.UpdateItem(context.TODO(), updateInput)
// 	if err != nil {
// 		fmt.Println("Error updating user item: ", err)
// 		return err
// 	}

// 	fmt.Println("Course added to user successfully.")
// 	return nil
// }
