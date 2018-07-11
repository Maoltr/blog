package controllers

import (
	"blog/model"
	"blog/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	product := services.CreateProduct(c.PostForm("name"), c.PostForm("description"), c.PostForm("images"),
		c.PostForm("price"))

	if validation(*product) {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Product item created successfully!",
			"resourceId": product.Id})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"message": "Product didn't create"})
}

func FetchAllProduct(c *gin.Context) {
	product := services.FetchAllProduct()

	if len(product) > 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": product})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
}

func FetchSingleProduct(c *gin.Context) {
	product := services.FetchSingleProduct(c.Param("id"))

	if product.Id != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": product})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
}

func UpdateProduct(c *gin.Context) {
	product := services.UpdateProduct(c.Param("id"), c.PostForm("name"), c.PostForm("description"),
		c.PostForm("images"), c.PostForm("price"))

	if product.Id != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": product})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})

}

func DeleteProduct(c *gin.Context) {
	res := services.DeleteProduct(c.Param("id"))

	if res {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product deleted successfully"})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})

}

func validation(product model.TransformedProduct) bool {
	if product.Id != 0 {
		return true
	}

	return false
}
