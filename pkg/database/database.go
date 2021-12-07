package database

import (
	"github.com/tonnytg/viewerlight/entity/entity"
	"log"
	"strconv"
)

var dataProduct []entity.Product

func init() {

	product1 := entity.Product{
		ID:          1,
		Name:        "Product 1",
		Price:       10.00,
		Description: "This is a product",
		Actions:     []string{"buy", "sell"},
	}
	product2 := entity.Product{
		ID:          2,
		Name:        "Product 2",
		Price:       20.00,
		Description: "This is a product",
		Actions:     []string{"buy", "sell"},
	}
	product3 := entity.Product{
		ID:          3,
		Name:        "Product 3",
		Price:       30.00,
		Description: "This is a product",
		Actions:     []string{"buy", "sell"},
	}
	dataProduct = append(dataProduct, product1, product2, product3)
}

func GetProducts() []entity.Product {
	return dataProduct
}

func  SaveProduct(p entity.Product) {
	dataProduct = append(dataProduct, p)
}

func UpdateProduct(p entity.Product) {

	for i, v := range dataProduct {
		if v.ID == p.ID {
			dataProduct[i] = entity.Product{
				ID:          p.ID,
				Name:        p.Name,
				Price:       p.Price,
				Description: p.Description,
				Actions:     []string{"save", "drop"},
			}
			break
		}
	}
}

func DeleteProduct(id string) {

	for i, v := range dataProduct {
		// convert srtring to int64
		idInt, _ := strconv.ParseInt(id, 10, 64)
		if v.ID == idInt {
			log.Println("will be deleted", idInt)
			dataProduct = append(dataProduct[:i], dataProduct[i+1:]...)
		}
	}
}
