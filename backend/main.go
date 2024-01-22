package main

import (
	"net/http"
	f "fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	. "github.com/duong-vriska/cvwo-assignment/backend/postHandler"
	. "github.com/duong-vriska/cvwo-assignment/backend/posts"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	
	database, err := gorm.Open("mysql", "root:nmai1202@tcp(localhost:3306)/db")
	if err != nil {
		f.Println("Error validating SQL arguments")
		panic(err.Error())
	}

	database.AutoMigrate(&Post{})

	f.Println("Successfully connected to SQL database")
	defer database.Close()
	http.ListenAndServe(":4000", setupServer())
}

func setupServer() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			return
		}
	})
	r.Mount("/posts", PostRoutes()) 
	return r
}

func PostRoutes() chi.Router {
	r := chi.NewRouter()
	postHandler := PostHandler{}
	r.Get("/", postHandler.ListPosts)
	r.Post("/new", postHandler.CreatePost)
	r.Get("/{id}", postHandler.GetPosts)
	r.Put("/{id}", postHandler.UpdatePost)
	r.Delete("/{id}", postHandler.DeletePost)
	// r.Get("/search/{category}", postHandler.FilterPost)
	return r
}

func browseHandler(w http.ResponseWriter, r *http.Request) {
	stmt:= "SELECT * FROM posts"
	rows, err:= db.Query(stmt)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var post Post
		err:= rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	stmt:= "SELECT * FROM posts WHERE id = ?"
	rows, err:= db.Query(stmt, id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var post Post
	for rows.Next() {
		err:= rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category)
		if err != nil {
			panic(err.Error())
		}
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {

	if title == "" | content == "" | category == "" {
		Println("Please fill in all fields")
		return
	}

	if r.Method == "POST" {
		id := GenerateRandomString(5)
		title := r.FormValue("title")
		content := r.FormValue("content")
		category := r.FormValue("category")
		stmt, err := db.Prepare("INSERT INTO `db`.`posts` (`post_id`, `title`, `content`, `category`) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		stmt.Exec(id, title, content, category)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	stmt, err := db.Prepare("DELETE FROM posts WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(id)
}

