package models

import (
	"CartShop/config"
	"CartShop/entities"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from product")
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	var product entities.Product
	for rows.Next() {
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.Photo)
		products = append(products, product)
	}
	return products, nil

}
func (*ProductModel) Find(id int64) (entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.Product{}, err
	}
	rows, err := db.Query("select * from product where id = ?", id)
	if err != nil {
		return entities.Product{}, err
	}

	var product entities.Product
	for rows.Next() {
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.Photo)
	}
	return product, nil

}
