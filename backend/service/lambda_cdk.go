package service

import (
	"encoding/json"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"

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

func DefineLambdas(stack *awscdk.Stack, usersTable awsdynamodb.Table, coursesTable awsdynamodb.Table, s3ImagesBucket awss3.Bucket, userPool awscognito.UserPool) {

	userService := NewUserHandler(*s3ImagesBucket.BucketName())
	userServiceJson, err := json.Marshal(userService)
	if err != nil {
		return
	}

	coursesService := NewCoursesHandler(*s3ImagesBucket.BucketName())
	serviceJSON, err := json.Marshal(coursesService)
	if err != nil {
		return
	}

	usersServiceJSON, err := json.Marshal(userService)
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
		CorsPreflight: &awscdkapigatewayv2alpha.CorsPreflightOptions{
			AllowOrigins: &[]*string{aws.String("*")},
			AllowMethods: &[]awscdkapigatewayv2alpha.CorsHttpMethod{
				awscdkapigatewayv2alpha.CorsHttpMethod_GET,
				awscdkapigatewayv2alpha.CorsHttpMethod_POST,
				awscdkapigatewayv2alpha.CorsHttpMethod_PUT,
				awscdkapigatewayv2alpha.CorsHttpMethod_DELETE,
			},
			AllowHeaders:  &[]*string{aws.String("Content-Type"), aws.String("Authorization")},
			ExposeHeaders: &[]*string{aws.String("Access-Control-Allow-Origin")},
		},
	})

	coursesFunctionIntg := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/courses"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_GET},
		Integration: coursesFunctionIntg})

	authorizer := awscdkapigatewayv2alpha.NewHttpAuthorizer(*stack, jsii.String("MyHttpAuthorizer"), &awscdkapigatewayv2alpha.HttpAuthorizerProps{
		AuthorizerName: jsii.String("MyHttpAuthorizer"),
		HttpApi:        api,
		Type:           awscdkapigatewayv2alpha.HttpAuthorizerType_JWT,
		JwtIssuer:      jsii.String("https://cognito-idp.eu-central-1.amazonaws.com/" + *userPool.UserPoolId()),
		JwtAudience:    jsii.Strings("3s9evb0dc0qspc503fnbnajgqm"), // Look this up in Cognito Userpool App Client settings. It’s the App client ID.
		IdentitySource: jsii.Strings("$request.header.Authorization"),
	})

	httpApiAuthorizer := awscdkapigatewayv2alpha.HttpAuthorizer_FromHttpAuthorizerAttributes(*stack, jsii.String("MyHttpAuthorizer4Test"), &awscdkapigatewayv2alpha.HttpAuthorizerAttributes{
		AuthorizerId:   authorizer.AuthorizerId(),
		AuthorizerType: jsii.String("JWT"),
	})

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

	// UPLOAD OBJECT
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

	s3ImagesBucket.GrantReadWrite(coursesFunction3, true)

	// CREATE COURSE
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

	// COURSE CONTRACT
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

	// ADD SECTION
	coursesFunction6 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("add_section"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"COURSES_SERVICE": aws.String(string(serviceJSON))},
			Entry:       jsii.String("../lambdas/add_section")})

	coursesFunctionIntg6 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction6, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/courses/section"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_POST},
		Integration: coursesFunctionIntg6})

	coursesTable.GrantReadWriteData(coursesFunction6)

	// ADD VIDEO
	coursesFunction7 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("add_video"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"COURSES_SERVICE": aws.String(string(serviceJSON))},
			Entry:       jsii.String("../lambdas/add_video")})

	coursesFunctionIntg7 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction7, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/courses/video"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_POST},
		Integration: coursesFunctionIntg7})

	coursesTable.GrantReadWriteData(coursesFunction7)

	// SIGN UP
	usersFunction := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("sign_up"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"USERS_SERVICE": aws.String(string(userServiceJson))},
			Entry:       jsii.String("../lambdas/sign_up")})

	usersFunction.AddToRolePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("cognito-idp:SignUp"),
			jsii.String("cognito-idp:AdminAddUserToGroup"),
		},
		Resources: &[]*string{userPool.UserPoolArn()},
	}))

	usersTable.GrantWriteData(usersFunction)

	usersFunctionIntg := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("users-function-integration"), usersFunction, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/users"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_POST},
		Integration: usersFunctionIntg,
	})

	// SIGN IN
	signInFunction := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("sign_in"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"USERS_SERVICE": aws.String(string(userServiceJson))},
			Entry:       jsii.String("../lambdas/sign_in")})

	signInFunction.AddToRolePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("cognito-idp:InitiateAuth"),
		},
		Resources: &[]*string{userPool.UserPoolArn()},
	}))

	usersTable.GrantWriteData(usersFunction)

	signInFunctionIntg := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("sign-in-function-integration"), signInFunction, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/users/login"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_POST},
		Integration: signInFunctionIntg,
	})

	// GET USER COURSES
	coursesFunction8 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("get_users_courses"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"USERS_SERVICE": aws.String(string(usersServiceJSON))},
			Entry:       jsii.String("../lambdas/get_users_courses")})

	coursesFunctionIntg8 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction8, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/users/courses"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_GET},
		Integration: coursesFunctionIntg8,
		Authorizer:  httpApiAuthorizer})

	usersTable.GrantReadWriteData(coursesFunction8)
	coursesTable.GrantReadWriteData(coursesFunction8)

	// Add USER COURSE
	coursesFunction9 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("add_users_course"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"USERS_SERVICE": aws.String(string(usersServiceJSON))},
			Entry:       jsii.String("../lambdas/add_users_course")})

	coursesFunctionIntg9 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction9, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/users/courses/{id}"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_PUT},
		Integration: coursesFunctionIntg9})

	usersTable.GrantReadWriteData(coursesFunction9)

	// GET OBJECT
	coursesFunction10 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("get_object"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"COURSES_SERVICE": aws.String(string(serviceJSON))},
			Entry:       jsii.String("../lambdas/get_object")})

	coursesFunctionIntg10 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction10, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/courses/get-object/{path}"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_GET},
		Integration: coursesFunctionIntg10})

	s3ImagesBucket.GrantRead(coursesFunction10, true)

	// Add USER COURSE
	coursesFunction11 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("get_user_by_username"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{"USERS_SERVICE": aws.String(string(usersServiceJSON))},
			Entry:       jsii.String("../lambdas/get_user_by_username")})

	coursesFunctionIntg11 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction11, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/users/username"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_GET},
		Integration: coursesFunctionIntg11})

	usersTable.GrantReadData(coursesFunction11)

	// Add USER COURSE
	time := *aws.Float64(60.0)
	coursesFunction12 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("add_watched_video"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Timeout:     awscdk.Duration_Seconds(&time),
			Environment: &map[string]*string{"USERS_SERVICE": aws.String(string(usersServiceJSON))},
			Entry:       jsii.String("../lambdas/add_watched_video")})

	coursesFunctionIntg12 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction12, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/users/watched"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_PUT},
		Integration: coursesFunctionIntg12})

	usersTable.GrantReadWriteData(coursesFunction12)
	s3ImagesBucket.GrantRead(coursesFunction12, true)
	coursesTable.GrantReadWriteData(coursesFunction12)

	coursesFunction13 := awscdklambdagoalpha.NewGoFunction(*stack, jsii.String("receive_certifikate"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime:     awslambda.Runtime_GO_1_X(),
			Timeout:     awscdk.Duration_Seconds(&time),
			Environment: &map[string]*string{"USERS_SERVICE": aws.String(string(usersServiceJSON))},
			Entry:       jsii.String("../lambdas/receive_certifikate")})

	coursesFunctionIntg13 := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(
		jsii.String("courses-function-integration"), coursesFunction13, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/users/certifikate/{id}"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_POST},
		Integration: coursesFunctionIntg13})

	usersTable.GrantReadWriteData(coursesFunction13)
	s3ImagesBucket.GrantRead(coursesFunction13, true)
	coursesTable.GrantReadWriteData(coursesFunction13)
}
