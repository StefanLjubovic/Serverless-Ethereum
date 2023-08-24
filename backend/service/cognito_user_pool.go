package service

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/jsii-runtime-go"
)

func CreateUserPool(stack *awscdk.Stack) *awscognito.UserPool {
	userPool := awscognito.NewUserPool(*stack, jsii.String("ethereum-user-pool"), &awscognito.UserPoolProps{
		UserPoolName:        jsii.String("ethereum-user-pool"),
		SignInCaseSensitive: jsii.Bool(false),
		UserVerification: &awscognito.UserVerificationConfig{
			EmailSubject: jsii.String("Verify your email"),
			EmailStyle:   awscognito.VerificationEmailStyle_CODE,
		},
		//SignInAliases: &awscognito.SignInAliases{
		//	Username: jsii.Bool(true),
		//	Email:    jsii.Bool(true),
		//},
		//AutoVerify: &awscognito.AutoVerifiedAttrs{
		//	Email: jsii.Bool(true),
		//},
	})

	return &userPool
}
