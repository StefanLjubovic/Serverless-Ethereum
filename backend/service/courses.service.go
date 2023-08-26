package service

import (
	"backend/dto"
	model "backend/model"
	"backend/repository"
	"backend/util"
	"context"
	"errors"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type CoursesService struct {
	DynamoDbClient   *dynamodb.Client
	S3ImageService   S3ImageServiceImpl
	CourseRepository repository.CoursesDynamoDBStore
	UsersReposiotry  repository.UsersDynamoDBStore
}

func NewCoursesHandler(bucketName string) *CoursesService {
	s3Service := NewS3ClientService(bucketName)
	usersRepository := repository.NewUsersDBStore("User")
	repository := repository.NewCoursesDBStore("Course")
	return &CoursesService{
		CourseRepository: *repository,
		DynamoDbClient:   nil,
		S3ImageService:   *s3Service,
		UsersReposiotry:  *usersRepository,
	}
}

func (coursesService *CoursesService) GetAllCourses() (*[]model.Course, error) {
	return coursesService.CourseRepository.GetAllCourses()
}

func (coursesService *CoursesService) GetCourseById(id string) (*model.Course, error) {
	return coursesService.CourseRepository.GetCourseById(id)
}

func (coursesService *CoursesService) UploadObject(image []byte) (string, error) {
	coursesService.S3ImageService.Start(context.TODO())
	return coursesService.S3ImageService.UploadObject(context.TODO(), image)
}

func (coursesService *CoursesService) Save(dto dto.CourseCreate) error {
	ethPrice, err := util.ConvertUSDToETH(dto.PriceUSD)
	if err != nil {
		return err
	}
	return coursesService.CourseRepository.Save(ethPrice, dto)
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

func (coursesService *CoursesService) AddSection(req dto.AddSectionDTO) error {
	course, err := coursesService.CourseRepository.GetCourseById(strconv.Itoa(req.ID))
	if err != nil {
		return err
	}
	section := model.Section{
		Name:   req.SectionName,
		Videos: []model.Video{},
	}
	course.Sections = append(course.Sections, section)
	return coursesService.CourseRepository.UpdateSections(course.Sections, course.ID)
}

func (coursesService *CoursesService) AddVideo(req dto.AddVideoDTO) error {
	course, err := coursesService.CourseRepository.GetCourseById(strconv.Itoa(req.CourseID))
	if err != nil {
		return err
	}
	var targetSection *model.Section
	for idx, section := range course.Sections {
		if section.Name == req.SectionName {
			targetSection = &course.Sections[idx]
			break
		}
	}
	if targetSection == nil {
		return errors.New("Section not found")
	}
	newVideo := model.Video{
		Name:    req.VideoName,
		Length:  req.Length,
		Path:    req.VideoPath,
		Watched: false,
	}
	targetSection.Videos = append(targetSection.Videos, newVideo)
	err = coursesService.CourseRepository.UpdateSections(course.Sections, req.CourseID)
	if err != nil {
		return err
	}

	return nil
}
