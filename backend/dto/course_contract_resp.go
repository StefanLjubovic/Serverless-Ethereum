package dto

import "math/big"

type CourseContractResp struct {
	PriceInWei *big.Int `json:"price_in_wei"`
	ID         uint     `json:"id"`
}
