package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Price       string `json:"price"`
}

func NewProduct(name, description, images, price string) *Product {
	product := Product{
		Name:        name,
		Description: description,
		Images:      images,
		Price:       price,
	}

	return &product
}
