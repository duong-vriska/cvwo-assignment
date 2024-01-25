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
        _, err := w.Write([]byte("Main page"))
        if err != nil {
            return
        }
    })

    return r
}