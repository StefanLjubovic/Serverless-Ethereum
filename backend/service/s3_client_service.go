package service

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

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

func (store *S3ImageServiceImpl) UploadObject(ctx context.Context, image []byte) (string, error) {
	filename := uuid.New()
	r := bytes.NewReader(image)
	_, err := store.s3session.PutObject(&s3.PutObjectInput{
		Body:   r,
		Bucket: aws.String(store.BUCKET_NAME),
		Key:    aws.String(filename.String()),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return filename.String(), nil
}

func (store *S3ImageServiceImpl) GetObject(ctx context.Context, filename string) ([]byte, error) {
	resp, err := store.s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(store.BUCKET_NAME),
		Key:    aws.String(filename),
	})

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
