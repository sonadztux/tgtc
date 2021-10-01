package service

import (
	"database/sql"

	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

func UpdateProduct(param dictionary.Product) (interface{}, error) {
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
	UPDATE
		products
	SET
		product_name = $1,
		product_price = $2,
		product_image = $3,
		shop_name = $4
	WHERE
		product_id = $5
	RETURNING
		product_id,
		product_name,
		product_price,
		product_image,
		shop_name
	`

	// actual query process
	row = db.QueryRow(query, param.Name, param.ProductPrice, param.ImageURL, param.ShopName, param.ID)
	err = row.Scan(&product.ID, &product.Name, &product.ProductPrice, &product.ImageURL, &product.ShopName)
	if err != nil {
		return product, err
	}

	return product, err
}
