package entity

type ProductRepository interface {
	Insert(id string, name string, price float64, errorMessage string) error
}
