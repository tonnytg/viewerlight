package process_product

import "github.com/tonnytg/viewerlight/entity/entity"

type ProcessProduct struct {
	Repository entity.ProductRepository
}

func NewProcessTransaction(repository entity.ProductRepository) *ProcessProduct {
	return &ProcessProduct{Repository: repository}
}

func (p *ProcessProduct) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	product := entity.NewProduct()
	product.Name = input.ProductId
	product.Price = input.Price
	product.Description = input.Description

	// check price business model
	if _, err := product.CheckStatus(); err != nil {
		return TransactionDtoOutput{}, err
	}

	return TransactionDtoOutput{
		ProductId:   product.Name,
		Price:       product.Price,
		Description: product.Description,
	}, nil
}
