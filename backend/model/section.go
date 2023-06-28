package model

type Section struct {
	Name  string `dynamodbav:"name"`
	Video Video  `dynamodbav:"video"`
}
