package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

func GetProducts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// variable declarations
	var (
		products []dictionary.Product
		data     dictionary.APIResponseProducts
	)
	// some input validations here

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
		log.Println(err)
	}
	defer rows.Close()

	// loop and struct scan
	for rows.Next() {
		var (
			p dictionary.Product
		)
		err = rows.Scan(&p.ID, &p.Name, &p.ProductPrice, &p.ImageURL, &p.ShopName)
		if err != nil {
			break
		}
		products = append(products, p)
	}

	// construct api response json
	if err != nil {
		data.Error = err.Error()
	} else {
		data.Products = products
	}

	resp, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	w.Write(resp)
}
