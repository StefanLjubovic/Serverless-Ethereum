package service

import (
	"encoding/json"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/jsii-runtime-go"
)

const getCoursesDir = "../service/lambdas"

func DefineLambdas(stack *awscdk.Stack, usersTable awsdynamodb.Table, coursesTable awsdynamodb.Table) {

	// userService := NewUserHandler(usersTable)

	coursesService := NewCoursesHandler()

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

	api := awscdkapigatewayv2alpha.NewHttpApi(*stack, jsii.String("http-api"), nil)

	coursesFunctionIntg := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(jsii.String("courses-function-integration"), coursesFunction, nil)

	api.AddRoutes(&awscdkapigatewayv2alpha.AddRoutesOptions{
		Path:        jsii.String("/courses"),
		Methods:     &[]awscdkapigatewayv2alpha.HttpMethod{awscdkapigatewayv2alpha.HttpMethod_GET},
		Integration: coursesFunctionIntg})

}
