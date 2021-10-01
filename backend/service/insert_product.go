package service

import (
	"database/sql"

	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

func InsertProduct(param dictionary.Product) (interface{}, error) {
	// variable declarations
	var (
		product dictionary.Product
		err     error
		row     *sql.Row
	)
	// get current database connection
	db := database.GetDB()

	// construct sql statement
	query := `
	INSERT INTO products
	(
		product_name,
		product_price,
		product_image,
		shop_name
	)
	VALUES(
		$1,
		$2,
		$3,
		$4
	)
	RETURNING
		product_id,
		product_name,
		product_price,
		product_image,
		shop_name
	`

	// actual query process
	row = db.QueryRow(query, param.Name, param.ProductPrice, param.ImageURL, param.ShopName)
	err = row.Scan(&product.ID, &product.Name, &product.ProductPrice, &product.ImageURL, &product.ShopName)
	if err != nil {
		return product, err
	}

	return product, err
}
