package service

import (
	"backend/dto"
	model "backend/model"
	"backend/repository"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UsersService struct {
	DynamoDbClient    *dynamodb.Client
	TableName         string
	UserRepository    repository.UsersDynamoDBStore
	CoursesRepository repository.CoursesDynamoDBStore
}

func NewUserHandler() *UsersService {
	coursesRepository := repository.NewCoursesDBStore("Course")
	repository := repository.NewUsersDBStore("User")
	return &UsersService{
		TableName:         "User",
		DynamoDbClient:    nil,
		UserRepository:    *repository,
		CoursesRepository: *coursesRepository,
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

func (usersService *UsersService) AddUsersCourse(username string, courseId int) error {

	usersCourse := model.UsersCourse{
		Course:          courseId,
		LastTimeWatched: time.Now(),
	}
	user, err := usersService.UserRepository.GetByUsername(username)
	if err != nil {
		return err
	}
	user.UsersCourses = append(user.UsersCourses, usersCourse)
	err = usersService.UserRepository.UpdateUsersCourses(user.ID, &user.UsersCourses)
	if err != nil {
		return err
	}
	return nil
}

func (usersService *UsersService) GetUsersCourses(username string) (*[]dto.CourseLastTimeWatched, error) {

	user, err := usersService.UserRepository.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	sortedUsersCourses := mergeSort(user.UsersCourses)
	courses := []dto.CourseLastTimeWatched{}
	for _, c := range sortedUsersCourses {
		course, err := usersService.CoursesRepository.GetCourseById(strconv.FormatUint(uint64(c.Course), 10))
		if err != nil {
			return nil, err
		}
		temp := dto.CourseLastTimeWatched{
			Course:          *course,
			LastTimeWatched: c.LastTimeWatched,
		}
		courses = append(courses, temp)
	}
	return &courses, nil
}

func mergeSort(usersCourses []model.UsersCourse) []model.UsersCourse {

	if len(usersCourses) < 2 {
		return usersCourses
	}
	first := mergeSort(usersCourses[:len(usersCourses)/2])
	second := mergeSort(usersCourses[len(usersCourses)/2:])
	return merge(first, second)
}

func merge(a []model.UsersCourse, b []model.UsersCourse) []model.UsersCourse {
	final := []model.UsersCourse{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i].LastTimeWatched.After(b[j].LastTimeWatched) {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
