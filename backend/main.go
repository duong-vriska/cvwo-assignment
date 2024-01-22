package main

import (
	"net/http"
	f "fmt"
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	. "github.com/duong-vriska/cvwo-assignment/backend/postHandler"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	db, err := sql.Open("mysql", "root:nmai1202@tcp(localhost:3306)/db")
	if err != nil {
		f.Println("Error validating SQL arguments")
		panic(err.Error())
	}

	err = db.Ping()
	if (err != nil) {
		f.Println("Error validating SQL connection")
		return 
	}

	insert,err := db.Query("INSERT INTO `db`.`post` (`id`, `title`, `content`) VALUES ('3', 'Hello World', 'This is a test post')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	f.Println("Successfully connected to SQL database")
	defer db.Close()
	http.ListenAndServe(":8080", setupServer())
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
	r.Get("/search/{category}", postHandler.FilterPost)
	return r
}
