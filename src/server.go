package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	// http.HandleFunc("/", getAllPosts)

	http.HandleFunc("/post", getPost)

	http.ListenAndServe(":3000", nil)
}

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

func loadPostsData() {
	jsonFile, err := os.Open("src/json/posts.json")

	if err != nil {
		fmt.Println("Error load file json", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &posts)
}
