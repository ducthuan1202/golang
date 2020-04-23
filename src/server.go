package main

import (
	"packages/header"
	qb "packages/querybuilder"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"os"
	"strconv"
)

type User struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post{}

func main() {

	loadPostsData()

	http.HandleFunc("/", getAllPosts)

	http.HandleFunc("/post", getPost)

	http.HandleFunc("/ads.js", actionAds)

	http.ListenAndServe(":3000", nil)
}
////////////////////////////////////////////////////////////////////////////////////////

func getPost(w http.ResponseWriter, r *http.Request) {

	// get the ID of the post from the route parameter
	var idParam string = r.FormValue("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	post := posts[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// parser 1 file sau đó execute
func actionAds(w http.ResponseWriter, r *http.Request)  {
	type Ads struct {
		Width	int
		Height 	int 
		Text 	string
	}
	
	defer showErr()
	// w.Header().Set("Content-Type", "application/javascript")	
	w.Header().Set(header.ContentType("javascript"))

	tmpl, err := template.ParseFiles("src/views/camp.tpl")
	checkErr(err)

	tbl := qb.Table("users")
	fmt.Println(tbl)

	var tool qb.QueryBuilder

	tool.
		Select("id", "name", "email").
		AddSelect("role", "status").
		AddSelect("   COUNT(id) AS 'total'    ").
		From("users").
		Where("id", ">", "10").
		AndWhere("status", "=", "1").
		OrWhere("role", "=", "admin")

	data := Ads {
		Width: 300,
		Height: 250,
		Text: fmt.Sprintf("%s", tool.GetQueryString()),
	}

	tmpl.Execute(w, data)
}

////////////////////////////////////////////////////////////////////////////////////////

func loadPostsData() {
	jsonFile, err := os.Open("src/json/posts.json")

	if err != nil {
		fmt.Println("Error load file json", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &posts)
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}

func showErr(){
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}