package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/duong-vriska/cvwo-assignment/postHandler"
	"github.com/duong-vriska/cvwo-assignment/posts"

	"net/http"
<<<<<<< HEAD
=======
	_ "myapp/postHandler"
	_ "myapp/posts"
>>>>>>> c5be727615b4bd40ca8d94778bcd2e1195f158cc
)

func main() {
	r := setupServer()
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		return
	}
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
	return r
}
