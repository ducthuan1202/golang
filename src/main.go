package main

import (
	"fmt"
	"net/http"
	"os"
	"packages/middleware"
	"text/template"
)

type NewsAggPage struct {
	Title string
	News  string
}

func (n NewsAggPage) String() string {
	return n.Title + " : " + n.News
}

func init() {

	// Set timezone
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

	// gán thư mục static là static
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// gán favicon.ico render static file
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/img/favicon.ico")
	})
}

func main() {

	// combine middleware with action handle
	finalHandler := http.HandlerFunc(actionHome)

	http.Handle("/", middleware.AddLogRaquest(finalHandler))

	http.ListenAndServe(":3000", nil)
}

// action
func actionHome(w http.ResponseWriter, r *http.Request) {

	p := NewsAggPage{Title: "Tieu De", News: "Tin noi bat"}

	t, err := template.ParseFiles("src/views/tmp.html")

	fmt.Println(t)

	if err != nil {
		fmt.Println(err)
		return
	}

	t.ExecuteTemplate(os.Stdout, "tmp", p)
	fmt.Println(t)
}
