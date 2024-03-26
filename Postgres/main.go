package main

import (
	"github.com/hakanaltindis/golang201/models"
)

func main() {

	product := models.Product{
		Title:       "Go Programming Language Book",
		Description: "This is a beautiful book",
		Price:       44.99,
	}
	models.InsertProduct(product)

	models.GetProducts()
}
