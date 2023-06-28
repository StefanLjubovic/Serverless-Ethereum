package model

type Price struct {
	PriceETH float64 `dynamodbav:"price_eth"`
	PriceUSD float64 `dynamodbav:"price_usd"`
}
