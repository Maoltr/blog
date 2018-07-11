package services

import (
	"blog/database"
	"blog/model"
)

func CreateProduct(name, description, images, price string) *model.TransformedProduct {
	var product model.Product
	product = *model.NewProduct(name, description, images, price)

	database.SaveProduct(&product)

	res := model.NewTransformedProduct(product)

	return res
}

func FetchAllProduct() []model.TransformedProduct {
	products := database.GetAllProducts()
	var result []model.TransformedProduct

	for _, item := range products {
		result = append(result, *model.NewTransformedProduct(item))
	}

	return result
}

func FetchSingleProduct(id string) *model.TransformedProduct {
	product := database.GetProductById(id)

	res := model.NewTransformedProduct(*product)

	return res
}

func UpdateProduct(id, name, description, images, price string) *model.TransformedProduct {
	product := database.UpdateProduct(id, name, description, images, price)

	res := model.NewTransformedProduct(*product)

	return res
}

func DeleteProduct(id string) bool {
	return database.DeleteProduct(id)
}
