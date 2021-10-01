package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/radityaqb/tgtc/backend/dictionary"
	"github.com/radityaqb/tgtc/backend/service"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// variable declarations
	var (
		resp  dictionary.APIResponse
		param dictionary.Product
		err   error
	)
	// some input validations here
	err = json.NewDecoder(r.Body).Decode(&param)
	// some input validations here
	//
	//
	//

	// input validated
	if err == nil {
		// proceed to the main service
		resp.Data, err = service.UpdateProduct(param)
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}

	w.Write(jsonResponse)
}
