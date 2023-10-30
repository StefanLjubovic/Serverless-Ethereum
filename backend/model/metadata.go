package model

type NFTMetadata struct {
	Name        NFTMetadataProperty `json:"name"`
	Description NFTMetadataProperty `json:"description"`
	Image       NFTMetadataProperty `json:"image"`
}

type NFTMetadataScheme struct {
	Title      string      `json:"title"`
	Type       string      `json:"type"`
	Properties NFTMetadata `json:"properties"`
}
type NFTMetadataProperty struct {
	Type        string      `json:"type"`
	Description interface{} `json:"description"`
}
