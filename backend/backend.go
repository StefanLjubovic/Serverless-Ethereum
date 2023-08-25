package main

import (
	service "backend/service"
	"os"

	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/joho/godotenv"
)

type BackendStackProps struct {
	awscdk.StackProps
}

func NewBackendStack(scope constructs.Construct, id string, props *BackendStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// service.DefineCognitoUserPool(&stack)
	service.DefineS3Bucket(&stack)
	usersTable := service.CreateUsersTable(&stack)
	coursesTable := service.CreateCourseTable(&stack)
	s3ImagesBucket := service.CreateCourseBucket(&stack)
	userPool := service.CreateUserPool(&stack)
	service.DefineLambdas(&stack, *usersTable, *coursesTable, *s3ImagesBucket, *userPool)
	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	NewBackendStack(app, "BackendStack", &BackendStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	account := os.Getenv("CDK_DEPLOY_ACCOUNT")
	region := os.Getenv("CDK_DEPLOY_REGION")
	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
	}
}
