package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/radityaqb/tgtc/backend/dictionary"
)

var (
	productMap map[int64]dictionary.Product = make(map[int64]dictionary.Product)
	idx = int64(1)
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	newProduct.ID = idx
	idx++
	productMap[newProduct.ID] = newProduct
	
	fmt.Fprintf(w, fmt.Sprint("success, id product: ", newProduct.ID))
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	idstring := r.URL.Query().Get("id")
	var product dictionary.Product

	idInt64, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	if _, productExist := productMap[idInt64]; productExist {
		product = productMap[idInt64]
	} else {
		http.Error(w, "product not found", 404)
		return 
	}

	productData, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(productData))
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if _, productExist := productMap[product.ID]; productExist {
		productMap[product.ID] = product
	} else {
		http.Error(w, "product not found", 404)
		return 
	}

	fmt.Fprintf(w, "success")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if _, productExist := productMap[product.ID]; productExist {
		delete(productMap, product.ID)
	} else {
		http.Error(w, "product not found", 404)
		return 
	}

	fmt.Fprintf(w, "success")
}