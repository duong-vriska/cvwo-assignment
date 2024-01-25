package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupServer() *chi.Mux {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"https://*", "http://*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
        MaxAge:           300, 
      }))

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        _, err := w.Write([]byte("Main page"))
        if err != nil {
            return
        }
    })

    return r
}