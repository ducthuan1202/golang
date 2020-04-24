package controller

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func IndexArticleHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, Articles")
}

func DetailArticleHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	article_id := vars["article_id"]
	fmt.Fprintf(w, "Article ID: %s", article_id)
}