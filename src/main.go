package main

import (
	"net/http"
	"packages/middleware"
	"os"
	"html/template"
	"fmt"
)

type NewsAggPage struct {
	Title string
	News string
}

func (n NewsAggPage) String()string{
	return n.Title + " : " + n.News
}

func init(){

	// Set timezone
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
}

func main() {
	finalHandler := http.HandlerFunc(actionHome)
  	http.Handle("/", middleware.AddLogRaquest(finalHandler))
  	http.ListenAndServe(":3000", nil)
}

// action 
func actionHome(w http.ResponseWriter, r *http.Request) {		

	p := NewsAggPage{ Title: "Tieu De", News: "Tin noi bat" }

	t, err := template.ParseFiles("src/views/tmp.html")	

	if err != nil {
		fmt.Println(err)
		return
	}

	t.ExecuteTemplate(os.Stdout, "tmp", p)

}
  