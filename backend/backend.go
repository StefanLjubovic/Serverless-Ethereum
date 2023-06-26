package main

import (
	"os"

	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/joho/godotenv"
)

type BackendStackProps struct {
	awscdk.StackProps
}

func DefineS3Bucket(stack *awscdk.Stack) {

	bucket := awss3.NewBucket(*stack, jsii.String("Serverless-Ethereum-frontend"), &awss3.BucketProps{
		WebsiteIndexDocument: jsii.String("index.html"),
		BlockPublicAccess: awss3.NewBlockPublicAccess(&awss3.BlockPublicAccessOptions{
			BlockPublicAcls:       jsii.Bool(false),
			BlockPublicPolicy:     jsii.Bool(false),
			IgnorePublicAcls:      jsii.Bool(false),
			RestrictPublicBuckets: jsii.Bool(false),
		}),
	})

	assetPath := "../frontend/build"
	awss3deployment.NewBucketDeployment(*stack, jsii.String("WebsiteDeployment"), &awss3deployment.BucketDeploymentProps{
		Sources: &[]awss3deployment.ISource{
			awss3deployment.Source_Asset(jsii.String(assetPath), nil),
		},
		DestinationBucket:    bucket,
		DestinationKeyPrefix: jsii.String("/"),
	})
	origin := awscloudfrontorigins.NewS3Origin(bucket, &awscloudfrontorigins.S3OriginProps{
		OriginPath: jsii.String("/"),
	})

	distribution := awscloudfront.NewDistribution(*stack, jsii.String("myDist"), &awscloudfront.DistributionProps{
		DefaultBehavior: &awscloudfront.BehaviorOptions{
			Origin: origin,
		},
	})

	bucketARN := *bucket.BucketArn() + "/*"

	bucket.AddToResourcePolicy(
		awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Actions: jsii.Strings("s3:GetObject"),
			Resources: jsii.Strings(
				bucketARN,
			),
			Effect: awsiam.Effect_ALLOW,
			Principals: &[]awsiam.IPrincipal{
				awsiam.NewAnyPrincipal(),
			},
		}),
	)

	// Output the CloudFront domain name
	awscdk.NewCfnOutput(*stack, jsii.String("DistributionDomainName"), &awscdk.CfnOutputProps{
		Value: distribution.DistributionDomainName(),
	})

	// api := awsapigateway.NewRestApi(stack, jsii.String("WebsiteApi"))
	// domainName := "<your_custom_domain_name>" // Replace with your custom domain name
	// domain := awsapigateway.NewDomainName(stack, jsii.String("WebsiteDomain"), &awsapigateway.DomainNameProps{
	// 	DomainName: jsii.String(domainName),
	// 	Certificate: awsapigateway.NewCertificate(stack, jsii.String("WebsiteCertificate"), &awsapigateway.CertificateProps{
	// 		DomainName:       jsii.String(domainName),
	// 		ValidationMethod: awsapigateway.ValidationMethod_DNS,
	// 	}),
	// })
	// api.DomainName().AddBasePathMapping(api.Root(), &awsapigateway.BasePathMappingOptions{
	// 	DomainName: domain,
	// })

	// // Create a Route53 hosted zone and map it to the custom domain
	// zone := awsroute53.NewHostedZone(stack, jsii.String("WebsiteHostedZone"), &awsroute53.HostedZoneProps{
	// 	ZoneName: jsii.String("<your_domain_name>"), // Replace with your domain name
	// })
	// awsroute53.NewARecord(stack, jsii.String("WebsiteARecord"), &awsroute53.ARecordProps{
	// 	Zone:       zone,
	// 	RecordName: jsii.String("<your_domain_name>"), // Replace with your domain name
	// 	Target:     awsroute53.RecordTarget_FromAlias(awsroute53targets.NewApiGateway(api)),
	// })

}

func NewBackendStack(scope constructs.Construct, id string, props *BackendStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	DefineS3Bucket(&stack)

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
