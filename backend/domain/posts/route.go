package posts

import (
	db "github.com/duong-vriska/cvwo-assignment/backend/database"

	"github.com/go-chi/chi/v5"
)

func RunRouter(db *db.Queries) *chi.Mux {
    r := chi.NewRouter()

    postHandler := NewHandler(db)

    r.Get("/", postHandler.HandleListPosts)
    r.Get("/{post_id}", postHandler.HandleGetPost)
    r.Post("/new", postHandler.HandleUpdatePost)
    r.Put("/{post_id}/edit", postHandler.HandleUpdatePost)
    r.Delete("/{id}", postHandler.HandleDeletePost)

    return r
}