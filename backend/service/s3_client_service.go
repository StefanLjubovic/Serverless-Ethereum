package service

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

var (
	REGION = "eu-central-1"
)

type S3ImageServiceImpl struct {
	BUCKET_NAME string
	s3session   *s3.S3
}

func NewS3ClientService(bucketName string) *S3ImageServiceImpl {
	s3Service := &S3ImageServiceImpl{
		BUCKET_NAME: bucketName,
	}
	return s3Service
}

func (store *S3ImageServiceImpl) Start(ctx context.Context) {
	mySession := session.Must(session.NewSession())
	store.s3session = s3.New(mySession, aws.NewConfig().WithRegion("eu-central-1"))
}

func (store *S3ImageServiceImpl) UploadObject(ctx context.Context, image []byte, ext string) (string, error) {
	filename := uuid.New().String() + "." + ext
	r := bytes.NewReader(image)
	_, err := store.s3session.PutObject(&s3.PutObjectInput{
		Body:   r,
		Bucket: aws.String(store.BUCKET_NAME),
		Key:    aws.String(filename),
	})
	path := "https://" + store.BUCKET_NAME + ".s3.eu-central-1.amazonaws.com" + "/" + filename
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return path, nil
}

func (store *S3ImageServiceImpl) GetObject(ctx context.Context, filename string) ([]byte, error) {
	parts := strings.Split(filename, "/")
	resp, err := store.s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(store.BUCKET_NAME),
		Key:    aws.String(parts[len(parts)-1]),
	})
	fmt.Println(parts[len(parts)-1])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
