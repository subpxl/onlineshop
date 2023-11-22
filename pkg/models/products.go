package models

import (
	"strings"
)

// func ProductList(w http.ResponseWriter, r *http.Request)
// // user can see products list
func GetAllProducts() []Product {
	var products []Product
	DB.Preload("Category").Find(&products)

	return products
}

// // user can see product page
func GetProductById(id int) Product {
	var product Product
	DB.Where("id = ?", id).First(&product)
	return product
}

func GetProductByName(name string) Product {
	var product Product
	DB.Where("link = ?", name).First(&product)
	return product
}

func CreateProduct(product Product) Product {
	link := strings.ReplaceAll(product.Name, " ", "-")
	product.Link = link
	DB.Create(&product)
	return product
}
func UpdateProduct(id int, product Product) Product {
	DB.Where("id= ? ", id).Updates(&product)
	return product
}
func DeleteProduct(id int) Product {
	var product Product
	DB.Where("id= ? ", id).First(&product)
	DB.Delete(product)
	return product
}
