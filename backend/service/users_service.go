package service

import (
	"backend/dto"
	model "backend/model"
	"backend/repository"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
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
	S3ImageService    S3ImageServiceImpl
	NFTService        NFTService
}

func NewUserHandler(bucketName string) *UsersService {
	coursesRepository := repository.NewCoursesDBStore("Course")
	s3Service := NewS3ClientService(bucketName)
	nftService := NewNFTHandler()
	repository := repository.NewUsersDBStore("User")
	return &UsersService{
		TableName:         "User",
		DynamoDbClient:    nil,
		UserRepository:    *repository,
		S3ImageService:    *s3Service,
		CoursesRepository: *coursesRepository,
		NFTService:        *nftService,
	}
}

func (usersService *UsersService) GetUserByUsername(username string) (*model.User, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		fmt.Println("Failed to make configuration: ", err)
		return nil, err
	}
	usersService.DynamoDbClient = dynamodb.NewFromConfig(cfg)
	params := &dynamodb.QueryInput{
		TableName:              aws.String("User"),
		IndexName:              aws.String("Username"),
		KeyConditionExpression: aws.String("username = :username"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":username": &types.AttributeValueMemberS{Value: username},
		},
	}

	queryResult, err := usersService.DynamoDbClient.Query(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	if len(queryResult.Items) == 0 {
		return nil, nil
	}
	user := &model.User{}
	err = attributevalue.UnmarshalMap(queryResult.Items[0], user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (usersService *UsersService) Save(dto dto.UserCreate) error {
	return usersService.UserRepository.Save(dto)
}

func (usersService *UsersService) AddUsersCourse(username string, courseId int) error {

	usersCourse := model.UsersCourse{
		Course:          courseId,
		LastTimeWatched: time.Now(),
		WatchedCount:    0,
		Watched:         make(map[string]bool),
	}
	user, err := usersService.UserRepository.GetByUsername(username)
	if err != nil {
		return err
	}
	if user.UsersCourses == nil {
		user.UsersCourses = []model.UsersCourse{}
	}
	user.UsersCourses = append(user.UsersCourses, usersCourse)
	err = usersService.UserRepository.UpdateUsersCourses(user.ID, &user.UsersCourses)
	if err != nil {
		return err
	}
	return nil
}

func (usersService *UsersService) AddCourse(username, id string) error {
	intValue, _ := strconv.Atoi(id)
	usersCourse := model.UsersCourse{
		Course:          intValue,
		LastTimeWatched: time.Now(),
		Watched:         make(map[string]bool),
		WatchedCount:    0,
	}
	return usersService.UserRepository.AddCourse(usersCourse, username)
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

func (usersService *UsersService) AddWatchedVideo(username string, dto *dto.AddWatchedDTO) error {
	user, err := usersService.UserRepository.GetByUsername(username)
	// if err != nil {
	// 	return err
	// }
	// var usersCourse model.UsersCourse
	// for _, course := range user.UsersCourses {
	// 	if course.Course == dto.CourseId {
	// 		if course.Watched[dto.Video] {
	// 			return nil
	// 		}
	// 		course.LastTimeWatched = time.Now()
	// 		course.Watched[dto.Video] = true
	// 		course.WatchedCount++
	// 		usersCourse = course
	// 	}
	// }
	// fmt.Println(usersCourse)
	// err = usersService.UserRepository.UpdateUsersCourses(user.ID, &user.UsersCourses)
	// if err != nil {
	// 	return err
	// }
	course, err := usersService.CoursesRepository.GetCourseById(strconv.Itoa(dto.CourseId))
	if err != nil {
		return err
	}
	usersService.S3ImageService.Start(context.TODO())
	img, _ := usersService.S3ImageService.GetObject(context.TODO(), course.Certificate.ImagePath)
	fmt.Println(img)
	return usersService.NFTService.ReceiveCertificate(user, course, img)
	// courseVideosSum := 0
	// for _, section := range course.Sections {
	// 	courseVideosSum += len(section.Videos)
	// }
	// if courseVideosSum-usersCourse.WatchedCount == 0 {
	// 	img, _ := usersService.S3ImageService.GetObject(context.TODO(), course.Certificate.ImagePath)
	// 	return receiveCertificate(user, course, img)
	// }
	// return nil
}
