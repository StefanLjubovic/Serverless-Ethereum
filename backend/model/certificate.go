package model

type Certificate struct {
	Name        string `dynamodbav:"name"`
	Description string `dynamodbav:"description"`
	ImagePath   string `dynamodbav:"image_path"`
}
