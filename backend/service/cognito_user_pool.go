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
			EmailStyle:   awscognito.VerificationEmailStyle_LINK,
		},
		AutoVerify: &awscognito.AutoVerifiedAttrs{
			Email: jsii.Bool(true),
		},
		SelfSignUpEnabled: jsii.Bool(true),
	})

	userPool.AddDomain(jsii.String("my-domain"), &awscognito.UserPoolDomainOptions{
		CognitoDomain: &awscognito.CognitoDomainOptions{
			DomainPrefix: jsii.String("ethereum-app"),
		},
	})

	userPool.AddClient(jsii.String("my-client"), &awscognito.UserPoolClientOptions{
		GenerateSecret: jsii.Bool(false),
		AuthFlows: &awscognito.AuthFlow{
			UserPassword: jsii.Bool(true),
			UserSrp:      jsii.Bool(true),
		},
	})

	return &userPool
}
