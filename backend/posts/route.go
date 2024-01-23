package posts

import (
	db "backend/database/schema"

	"github.com/go-chi/chi/v5"
	"github.com/duong-vriska/cvwo-assignment/backend/utils"
)

func InitRouter(db *db.Queries) *chi.Mux {
    r := chi.NewRouter()

    // initialise post handler
    h := NewHandler(db)

    // add routes to the subrouter
    r.Get("/", h.HandleGetAllPosts)
    r.Get("/threads/{thread_id}", h.HandleListPostsFromThread)
    r.Post("/new", h.HandleCreatePost)
    r.Put("/edit", h.HandleEditPost)
    r.Delete("/delete/{id}", h.HandleDeletePost)

    return r
}