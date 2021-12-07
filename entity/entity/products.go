package entity

import (
	"errors"
)

type Product struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Actions     []string `json:"actions"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) CheckStatus() (string, error) {
	if p.Price > 0 && p.Price < 1000 {
		return "success", nil
	}
	return "failed", errors.New("price must be 1 between 1000")
}
