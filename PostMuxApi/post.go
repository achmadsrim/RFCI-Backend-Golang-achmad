package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Author   string `json:"author"`
	Title    string `json:"title"`
	Comments []struct {
		Message string `json:"message"`
	} `json:"comments"`
}

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}

func main() {
	router := mux.NewRouter()
	posts = append(posts, Post{Author: "Who", Title: "My first post"})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	http.ListenAndServe(":8000", router)
}
