package dto

type CourseCreate struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	PriceUSD    float64 `json:"price_usd"`
	Image       string  `json:"image"`
	Certificate CertDTO `json:"certificate"`
}

type CertDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
}
