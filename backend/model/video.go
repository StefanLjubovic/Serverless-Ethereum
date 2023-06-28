package model

type Video struct {
	Name    string `dynamodbav:"name"`
	Length  int    `dynamodbav:"length"`
	Path    string `dynamodbav:"path"`
	Watched bool   `dynamodbav:"watched"`
}
