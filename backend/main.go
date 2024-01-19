package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := setupServer()
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}

func setupServer() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
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
	r.Post("/", postHandler.CreatePost)
	r.Get("/{id}", postHandler.GetPosts)
	r.Put("/{id}", postHandler.UpdatePost)
	r.Delete("/{id}", postHandler.DeletePost)
	return r
}
