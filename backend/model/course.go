package model

type Course struct {
	ID          int         `dynamodbav:"id"`
	Name        string      `dynamodbav:"name"`
	Description string      `dynamodbav:"description"`
	Sections    []Section   `dynamodbav:"sections"`
	Ratings     []Rating    `dynamodbav:"ratings"`
	Categories  []Category  `dynamodbav:"categories"`
	Price       Price       `dynamodbav:"price"`
	Image       string      `dynamodbav:"image"`
	Certificate Certificate `dynamodbav:"certificate"`
}

type Certificate struct {
	Name        string `dynamodbav:"name"`
	Description string `dynamodbav:"description"`
	ImagePath   string `dynamodbav:"image_path"`
}

type Rating struct {
	Grade   uint8  `dynamodbav:"grade"`
	Comment string `dynamodbav:"comment"`
}

type Section struct {
	Name   string  `dynamodbav:"name"`
	Videos []Video `dynamodbav:"videos"`
}

type Category int

const (
	DATASCIENCE Category = iota
	CYBERSECURITY
	MACHINELEARNING
	CLOUDCOMPUTING
	DESIGN
)
