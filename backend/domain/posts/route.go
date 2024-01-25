package posts

import (
    "net/http"

	"github.com/go-chi/chi/v5"
)

func PostRouter(pHandler *Handler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.HandleListPosts)
    r.Get("/{post_id}", pHandler.HandleGetPost)
    r.Get("/category/{category}", pHandler.HandleGetPostByCategory)
    r.Post("/new", pHandler.HandleCreatePost)
    r.Put("/{post_id}", pHandler.HandleUpdatePost)
    r.Delete("/{post_id}", pHandler.HandleDeletePost)
	return r
}