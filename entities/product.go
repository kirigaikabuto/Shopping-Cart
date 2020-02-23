package entities

type Product struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	Description string  `json:"description"`
	Photo       string  `json:"photo"`
}
