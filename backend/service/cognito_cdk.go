package service

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/jsii-runtime-go"
)

func DefineCognitoUserPool(stack *awscdk.Stack) {
	userPool := awscognito.NewUserPool(*stack, jsii.String("serverless-ethereum-user-pool"), &awscognito.UserPoolProps{
		UserPoolName:        jsii.String("serverless-ethereum-user-pool"),
		SignInCaseSensitive: jsii.Bool(false),
		UserVerification: &awscognito.UserVerificationConfig{
			EmailSubject: jsii.String("Verify your email"),
			EmailStyle:   awscognito.VerificationEmailStyle_CODE,
		},
		SignInAliases: &awscognito.SignInAliases{
			Username: jsii.Bool(true),
			Email:    jsii.Bool(true),
		},
		AutoVerify: &awscognito.AutoVerifiedAttrs{
			Email: jsii.Bool(true),
		},
	})

	fmt.Print(userPool.PhysicalName())
}
