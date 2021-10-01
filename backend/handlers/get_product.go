package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/radityaqb/tgtc/backend/dictionary"
	"github.com/radityaqb/tgtc/backend/service"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// variable declarations
	var (
		resp      dictionary.APIResponse
		err       error
		productID int
	)
	// some input validations here
	//
	//
	//

	r.ParseForm()
	vars := mux.Vars(r)
	productID, err = strconv.Atoi(vars["id"])

	// input validated
	if err == nil {
		// proceed to the main service
		resp.Data, err = service.GetProduct(productID)
	}

	// construct api response json
	if err != nil {
		resp.Error = err.Error()
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}

	w.Write(jsonResponse)
}
