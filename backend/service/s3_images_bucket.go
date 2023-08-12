package service

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/jsii-runtime-go"
)

func CreateCourseBucket(stack *awscdk.Stack) *awss3.Bucket {

	bucket := awss3.NewBucket(*stack, aws.String("CourseBucket"), &awss3.BucketProps{
		Versioned:     jsii.Bool(true),
		AccessControl: awss3.BucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
	})
	return &bucket

}
