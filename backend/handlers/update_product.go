package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// variable declarations
	var (
		data        dictionary.APIResponseSingleProduct
		row         *sql.Row
		param       dictionary.Product
		p           dictionary.Product
		err         error
		errorString string
	)
	// some input validations here
	if err = json.NewDecoder(r.Body).Decode(&param); err != nil {
		errorString = "Body yang dikirim tidak valid"
	}

	// lolos validasi
	if err == nil {
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
		err = row.Scan(&p.ID, &p.Name, &p.ProductPrice, &p.ImageURL, &p.ShopName)
		if err != nil {
			errorString = err.Error()
		}
	}

	data.Product = p
	data.Error = errorString

	resp, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	w.Write(resp)
}
