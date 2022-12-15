package main

import (
	"product/config"
	"product/entities"
	"fmt"
	"product/models"
)

func main() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {

		productModel := models.ProductModel{
			Db: db,
		}

		product := entities.Product{
			Id:       3,
			Name:     "def",
			Price:    888,
			Quantity: 999,
			Status:   false,
		}
		rowsAffected, err2 := productModel.Update(&product)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Rows Affected:", rowsAffected)
		}

	}
}
