package dto

import "github.com/joaocsv/hexagonal-example/app"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProductDTO() *Product {
	return &Product{}
}

func (p *Product) Bind(product *app.Product) (*app.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()

	if err != nil {
		return &app.Product{}, err
	}

	return product, nil
}
