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
