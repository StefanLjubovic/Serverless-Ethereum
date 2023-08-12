package dto

type CourseContract struct {
	SenderAddress string  `json:"sender_address"`
	PriceUSD      float64 `json:"price_usd"`
}
