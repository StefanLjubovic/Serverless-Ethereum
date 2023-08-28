package service

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

func CreateUsersTable(stack *awscdk.Stack) *awsdynamodb.Table {

	userTable := awsdynamodb.NewTable(*stack, jsii.String("User"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_NUMBER,
		},
		TableName:           jsii.String("User"),
		ReadCapacity:        jsii.Number(1),
		WriteCapacity:       jsii.Number(1),
		PointInTimeRecovery: jsii.Bool(true),
	})

	userTable.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String("Username"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})
	userTable.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String("Email"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("email"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	return &userTable

}

func CreateCourseTable(stack *awscdk.Stack) *awsdynamodb.Table {
	courseTable := awsdynamodb.NewTable(*stack, jsii.String("Course"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_NUMBER,
		},
		TableName:           jsii.String("Course"),
		ReadCapacity:        jsii.Number(2),
		WriteCapacity:       jsii.Number(1),
		PointInTimeRecovery: jsii.Bool(true),
	})
	courseTable.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String("SectionName"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("section.name"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})
	courseTable.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String("CourseName"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("name"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})
	courseTable.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String("Rating"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("rating.grade"),
			Type: awsdynamodb.AttributeType_NUMBER,
		},
	})
	return &courseTable
}
