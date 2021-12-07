package database

import (
	"github.com/tonnytg/viewerlight/entity/entity"
	"log"
	"strconv"
)

type DataValues struct {
	SliceProducts []entity.Product
}

func (dv *DataValues) StartInitialValues() {

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

	dv.SliceProducts = append(dv.SliceProducts, product1, product2)
}

func (dv *DataValues) GetProducts() []entity.Product {
	return dv.SliceProducts
}

func (dv *DataValues) SaveProduct() {
	dv.SliceProducts = append(dv.SliceProducts, entity.Product{})
}

func (dv *DataValues) UpdateProduct(p entity.Product) {

	for i, v := range dv.SliceProducts {
		if v.ID == p.ID {
			dv.SliceProducts[i] = entity.Product{
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

func (dv *DataValues) DeleteProduct(id string) {

	dv.SliceProducts = append(dv.SliceProducts, entity.Product{})

	for i, v := range dv.SliceProducts {
		// convert srtring to int64
		idInt, _ := strconv.ParseInt(id, 10, 64)
		if v.ID == idInt {
			log.Println("will be deleted", idInt)
			dv.SliceProducts = append(dv.SliceProducts[:i], dv.SliceProducts[i+1:]...)
		}
	}

}
