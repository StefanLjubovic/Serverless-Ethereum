package model

type Rating struct {
	Grade   uint8  `dynamodbav:"grade"`
	Comment string `dynamodbav:"comment"`
}
