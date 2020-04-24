package controller

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
)

type Product struct {
	Id 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Price 	float64 `json:"price"`
	Status 	string 	`json:"status"`
}

var products []Product = []Product{
	Product{ Id: 1, Name: "SP 01", Price: 20.5, Status: "New"},
	Product{ Id: 2, Name: "SP 02", Price: 20.5, Status: "Saleof"},
	Product{ Id: 3, Name: "SP 03", Price: 20.5, Status: "Feature"},
}

func IndexProductHandler(w http.ResponseWriter, r *http.Request)  {

	json, err := json.Marshal(products)
	if err != nil {
		fmt.Fprintf(w, "500 - Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func DetailProductHandler(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)	
	product_id, err := strconv.Atoi(vars["product_id"])

	if product_id < 0 || product_id >= len(products) || err != nil {
		fmt.Fprintf(w, "Product ID: %v not found", product_id)
		return
	}

	product := products[product_id]
	
	json, err := json.Marshal(product)
	if err != nil {
		fmt.Fprintf(w, "500 - Server Error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func AddProductHandler(w http.ResponseWriter, r *http.Request)  {

	countProducts := len(products)
	product := Product{ Id: countProducts + 1, Name: fmt.Sprintf("SP 0%v", countProducts), Price: 12, Status: "New"}
	products = append(products, product)	
	
	json, err := json.Marshal(product)
	if err != nil {
		fmt.Fprintf(w, "500 - Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
