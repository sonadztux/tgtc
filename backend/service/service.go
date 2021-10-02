package service

import (
	"fmt"

	"github.com/radityaqb/tgtc/backend/database"
)

func SampleFunction() {
	fmt.Printf("My Service!")

	// // you can connect and
	// // get current database connection
	db := database.GetDB()

	// // construct query
	// query := `
	// SELECT something FROM table_something WHERE id = $1
	// `
	// // actual query process
	// row = db.QueryRow(query, paramID)

	// // read query result, and assign to variable(s)
	// err = row.Scan(&ID, &name)
}
