package entities

type Item struct {
	Product  Product `json:"id"`
	Quantity int64   `json:"quantity"`
}
