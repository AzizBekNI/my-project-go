package models

import (
	"database/sql"
	"product/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) Update(product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("update product set name = ?, price = ?, quantity = ?, status = ? where id = ?", product.Name, product.Price, product.Quantity, product.Status, product.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}