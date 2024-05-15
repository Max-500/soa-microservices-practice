package entities

type Product struct {
	Uuid string   `json:"uuid"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type Products struct {
	Products []Product `json:"products"`
}
