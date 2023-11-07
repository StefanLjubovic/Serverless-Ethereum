package service

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
)

func DefineS3Bucket(stack *awscdk.Stack) {

	bucket := awss3.NewBucket(*stack, jsii.String("Serverless-Ethereum-frontend"), &awss3.BucketProps{
		WebsiteIndexDocument: jsii.String("index.html"),
		// BlockPublicAccess: awss3.NewBlockPublicAccess(&awss3.BlockPublicAccessOptions{
		// 	BlockPublicAcls:       jsii.Bool(false),
		// 	BlockPublicPolicy:     jsii.Bool(false),
		// 	IgnorePublicAcls:      jsii.Bool(false),
		// 	RestrictPublicBuckets: jsii.Bool(false),
		// }),
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

	account := os.Getenv("CDK_DEPLOY_ACCOUNT") // Assuming you've set the environment variable
	bucketARN := *bucket.BucketArn() + "/*"

	bucket.AddToResourcePolicy(
		awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Actions:   jsii.Strings("s3:GetObject"),
			Resources: jsii.Strings(bucketARN),
			Effect:    awsiam.Effect_ALLOW,
			Principals: &[]awsiam.IPrincipal{
				awsiam.NewServicePrincipal(aws.String("cloudfront.amazonaws.com"), nil),
			},
			Conditions: &map[string]interface{}{
				"StringEquals": &map[string]interface{}{
					"AWS:SourceArn": "arn:aws:cloudfront:" + account + ":distribution/" + *distribution.DistributionId(),
				},
			},
		}),
	)

	// Output the CloudFront domain name
	awscdk.NewCfnOutput(*stack, jsii.String("DistributionDomainName"), &awscdk.CfnOutputProps{
		Value: distribution.DistributionDomainName(),
	})

}
