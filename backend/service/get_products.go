package service

import (
	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

func GetProducts() (interface{}, error) {
	// variable declarations
	var (
		products []dictionary.Product
		err      error
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
	`

	// actual query process
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// loop and struct scan
	for rows.Next() {
		var (
			p dictionary.Product
		)
		err = rows.Scan(&p.ID, &p.Name, &p.ProductPrice, &p.ImageURL, &p.ShopName)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, err
}
