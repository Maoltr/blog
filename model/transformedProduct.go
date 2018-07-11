package model

type TransformedProduct struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Price       string `json:"price"`
}

func NewTransformedProduct(item Product) *TransformedProduct {
	res := TransformedProduct{
		Id:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Images:      item.Images,
		Price:       item.Price,
	}

	return &res
}
