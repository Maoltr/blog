package database

import "blog/model"

func SaveProduct(product *model.Product) {
	db := Database()
	defer Close(db)
	db.Save(&product)
}

func GetAllProducts() []model.Product {
	var products []model.Product

	db := Database()
	defer Close(db)
	db.Find(&products)

	return products
}

func GetProductById(id string) *model.Product {
	var product model.Product

	db := Database()
	defer Close(db)
	db.First(&product, id)

	return &product
}

func UpdateProduct(id, name, description, images, price string) *model.Product {
	var product model.Product

	db := Database()
	defer Close(db)
	db.First(&product, id)

	if product.ID != 0 {
		db.Model(&product).Update("name", name)
		db.Model(&product).Update("description", description)
		db.Model(&product).Update("images", images)
		db.Model(&product).Update("price", price)
	}

	return &product
}

func DeleteProduct(id string) bool {
	var product model.Product

	db := Database()
	defer Close(db)
	db.First(&product, id)

	if product.ID == 0 {
		return false
	}

	db.Delete(&product)

	return true
}
