package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ImageSrc  string `json:"image_src"`
}

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "blog_db",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If "OPTIONS" request, respond immediately
		if r.Header.Get("Access-Control-Request-Method") != "" {
			enableCors(&w)
			return
		}
	})
	router.GET("/posts", getPosts)
	router.POST("/posts", createPost)
	router.GET("/posts/:id", getPost)
	router.PUT("/posts/:id", updatePost)
	router.DELETE("/posts/:id", deletePost)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)
	var posts []Post
	result, err := db.Query("SELECT id, title, content, created_at,updated_at, image_src from posts")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.ImageSrc)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	stmt, err := db.Prepare("INSERT INTO posts(Title,Content,image_src) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	res, err := stmt.Exec(post.Title, post.Content, post.ImageSrc)
	if err != nil {
		panic(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	post.ID = int(id)
	json.NewEncoder(w).Encode(post)
}

func getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enableCors(&w)
	id, _ := strconv.Atoi(ps.ByName("id"))
	result, err := db.Query("SELECT id, title, content, created_at, updated_at, image_src FROM posts WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post Post
	for result.Next() {
		err := result.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.ImageSrc)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enableCors(&w)
	id, _ := strconv.Atoi(ps.ByName("id"))
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	stmt, err := db.Prepare("UPDATE posts SET title = ?, content = ?, image_src= ?, WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(post.Title, post.Content, id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %d was updated", id)
}

func deletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enableCors(&w)
	id, _ := strconv.Atoi(ps.ByName("id"))
	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %d was deleted", id)
}
