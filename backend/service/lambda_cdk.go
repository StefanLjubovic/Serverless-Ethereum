package service

import (
	"encoding/json"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/jsii-runtime-go"
)

const getCoursesDir = "../lambdas"

func DefineLambdas(stack *awscdk.Stack, usersTable awsdynamodb.Table, coursesTable awsdynamodb.Table, s3_images_bucket awss3.Bucket) {

	// userService := NewUserHandler(usersTable)

	coursesService := NewCoursesHandler(*s3_images_bucket.BucketName())

	serviceJSON, err := json.Marshal(coursesService)
	if err != nil {
		return
	}

	coursesFunction := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("get_all_courses"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"COURSES_SERVICE": aws.String(string(serviceJSON))},
			Entry:       jsii.String(getCoursesDir)})

	coursesTable.GrantReadData(coursesFunction)

	api := awscdkapigatewayv2alpha.NewHttpApi(*stack, jsii.String("http-api"), &awscdkapigatewayv2alpha.HttpApiProps{
		CorsPreflight: &awscdkapigatewayv2alpha.CorsPreflightOptions{AllowOrigins: &[]*string{aws.String("*")}},
	})

	coursesFunctionIntg := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/courses"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_GET},
		Integration: coursesFunctionIntg})

	// GET COURSE BY ID
	coursesFunction2 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("get_course_by_id"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"COURSES_SERVICE": aws.String(string(serviceJSON))},
			Entry:       jsii.String("../lambdas/get_course_by_id")})

	coursesTable.GrantReadData(coursesFunction2)

	coursesFunctionIntg2 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction2, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/courses/{id}"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_GET},
		Integration: coursesFunctionIntg2})

	coursesFunction3 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("upload_object"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"COURSES_SERVICE": aws.String(string(serviceJSON))},
			Entry:       jsii.String("../lambdas/upload_object")})

	coursesFunctionIntg3 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction3, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/courses/upload-object"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_POST},
		Integration: coursesFunctionIntg3})

	s3_images_bucket.GrantReadWrite(coursesFunction3, true)

	coursesFunction4 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("create_course"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"COURSES_SERVICE": aws.String(string(serviceJSON))},
			Entry:       jsii.String("../lambdas/create_course")})

	coursesFunctionIntg4 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction4, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/courses"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_POST},
		Integration: coursesFunctionIntg4})

	coursesTable.GrantWriteData(coursesFunction4)

	coursesFunction5 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("course_contract"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"COURSES_SERVICE": aws.String(string(serviceJSON))},
			Entry:       jsii.String("../lambdas/course_contract")})

	coursesFunctionIntg5 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction5, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/courses/contract/{price}"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_GET},
		Integration: coursesFunctionIntg5})
}
