package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/radityaqb/tgtc/backend/dictionary"
	"github.com/radityaqb/tgtc/backend/service"
)

func GetProducts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// variable declarations
	var (
		resp dictionary.APIResponse
		err  error
	)
	// some input validations here
	//
	//
	//

	// proceed to the main service
	resp.Data, err = service.GetProducts()

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
