package dto

import "backend/model"

type CourseCreate struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	PriceUSD    float64           `json:"price_usd"`
	Image       string            `json:"image"`
	Certificate model.Certificate `json:"certificate"`
}
