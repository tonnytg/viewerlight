package process_product

type TransactionDtoInput struct {
	Name        string
	Price       float64
	ProductId   string
	Description string
	Quantity    int
}

type TransactionDtoOutput struct {
	Name        string
	Price       float64
	ProductId   string
	Description string
	Quantity    int
}
