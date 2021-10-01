package service

import (
	"database/sql"

	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

func GetProduct(productID int) (interface{}, error) {
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
	SELECT
		product_id,
		product_name,
		product_price,
		product_image,
		shop_name
	FROM
		products
	WHERE
		product_id = $1
	`

	// actual query process
	row = db.QueryRow(query, productID)
	err = row.Scan(&product.ID, &product.Name, &product.ProductPrice, &product.ImageURL, &product.ShopName)
	if err != nil && err != sql.ErrNoRows {
		return product, err
	}

	return product, err
}
