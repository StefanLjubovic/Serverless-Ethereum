package model

type Section struct {
	Name   string  `dynamodbav:"name"`
	Videos []Video `dynamodbav:"videos"`
}
