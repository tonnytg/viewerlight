package products

type Product struct {
	ID          int64  `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       int64  `json:"price"`
}

func Create(p Product) (Product, error) {
    return p, nil
}