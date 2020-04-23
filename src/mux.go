package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main()  {

	// tạo route tự động 
	routes := map[string]http.HandlerFunc {
		"/": handler,
		"/products": productsHandler,
		"/articles": articlesHandler,
	}

	r := mux.NewRouter()

	for key, val := range routes {		
		r.HandleFunc(key, val).Methods("GET")
	}

	// route thủ công
	
	// r.HandleFunc("/", handler).Methods("GET")
	// r.HandleFunc("/products", productsHandler).Methods("GET")
	// r.HandleFunc("/articles", articlesHandler).Methods("GET").Queries("key", "value")

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, how are you?")
}

func productsHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, Products")
}

func articlesHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, Articles")
}