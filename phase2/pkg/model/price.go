package model

type PriceReq struct {
	Name string `json:"name"`
}
type RetrieverAssets struct {
	AssetID  string  `json:"asset_id"`
	Name     string  `json:"name"`
	PriceUSD float64 `json:"price_usd"`
}
