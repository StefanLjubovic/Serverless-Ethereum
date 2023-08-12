package dto

type CourseCreate struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	PriceUSD    float64 `json:"price_usd"`
	Image       string  `dynamodbav:"image"`
	Certificate string  `dynamodbav:"certificate"`
}
