package services

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/jsii-runtime-go"
)

func defineS3Bucket(stack awscdk.Stack) {

	bucket := awss3.NewBucket(stack, jsii.String("WebsiteBucket"), &awss3.BucketProps{
		WebsiteIndexDocument: jsii.String("index.html"),
		PublicReadAccess:     jsii.Bool(true),
	})

	assetPath := "../../frontend/build" // Replace with the path to the directory containing your website files
	awss3deployment.NewBucketDeployment(stack, jsii.String("WebsiteDeployment"), &awss3deployment.BucketDeploymentProps{
		Sources: &[]awss3deployment.ISource{
			awss3deployment.Source_Asset(jsii.String(assetPath), nil),
		},
		DestinationBucket:    bucket,
		DestinationKeyPrefix: jsii.String("web/static"),
	})

	resource := *bucket.BucketArn() + "/*"

	bucket.AddToResourcePolicy(
		awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Actions:   jsii.Strings("s3:GetObject"),
			Resources: jsii.Strings(resource),
			Effect:    awsiam.Effect_ALLOW,
		}))

	// bucket.AddToResourcePolicy(policy)

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
