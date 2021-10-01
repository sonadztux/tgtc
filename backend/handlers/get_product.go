package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// variable declarations
	var (
		data        dictionary.APIResponseSingleProduct
		row         *sql.Row
		p           dictionary.Product
		err         error
		errorString string
		productID   int
	)
	// some input validations here
	r.ParseForm()
	vars := mux.Vars(r)
	if productID, err = strconv.Atoi(vars["id"]); err != nil {
		errorString = "Parameter ID tidak valid"
	}

	// lolos validasi
	if err == nil {
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
		err = row.Scan(&p.ID, &p.Name, &p.ProductPrice, &p.ImageURL, &p.ShopName)
		if err != nil {
			if err == sql.ErrNoRows {
				errorString = "Data tidak ditemukan"
			} else {
				errorString = err.Error()
			}
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
