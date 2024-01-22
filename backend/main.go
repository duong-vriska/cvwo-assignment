package main

import (
	db "github.com/duong-vriska/cvwo-assignment/backend/database"
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

func PostRoutes(db *db.Queries) *chi.Mux {
	r := chi.NewRouter()
	postHandler := PostHandler{}
	r.Get("/", postHandler.ListPosts)
	r.Post("/new", postHandler.CreatePost)
	r.Get("/{id}", postHandler.GetPosts)
	r.Put("/{id}", postHandler.UpdatePost)
	r.Delete("/{id}", postHandler.DeletePost)
	r.Get("/search/{category}", postHandler.FilterPost)
	return r
}