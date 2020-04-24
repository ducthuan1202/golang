package controller

import (
	"fmt"
	"net/http"
)

func HomepageHandler(w http.ResponseWriter, r *http.Request)  {

	defer func ()  {
		if r := recover(); r != nil {
			errStr := fmt.Sprintf("%v", r)
			fmt.Fprintf(w, errStr)
			return
		}
	}()


	// urlProductsIndex, _ := rx.Get("products.index").URL("", "")
	// urlProductsDetail, _ := rx.Get("products.detail").URL("product_id", "1")


	// productIndexUrl, _ := r.Get("products.index").URL()	
	productIndexUrl := "/ahihi"

	fmt.Fprintf(w, "Home page. URL Products: %s", productIndexUrl)
}