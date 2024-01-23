package posts 

import (
	"net/http"

	db "github.com/duong-vriska/cvwo-assignment/backend/database"
	"github.com/duong-vriska/cvwo-assignment/backend/utils"
	"github.com/go-chi/chi/v5"

)

type Handler struct {
	db *db.Queries 
}

func NewHandler(db *db.Queries) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) HandleListPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := h.db.ListPost(ctx)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}
	utils.Respond(w, http.StatusOK, posts)
}

func (h *Handler) HandleGetPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := chi.URLParam(r, "post_ID")
	post, err := h.db.GetPost(ctx, postID)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}
	utils.Respond(w, http.StatusOK, post)
}

func (h *Handler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var post db.CreatePostParams
	utils.Parse(w, r, &post)
	createdPost, err := h.db.CreatePost(ctx, post)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}
	utils.Respond(w, http.StatusOK, createdPost)
}

func (h *Handler) HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var post db.UpdatePostParams
	utils.Parse(w, r, &post)
	updatedPost, err := h.db.UpdatePost(ctx, post)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}
	utils.Respond(w, http.StatusOK, updatedPost)
}

func (h *Handler) HandleDeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := chi.URLParam(r, "post_ID")
	err := h.db.DeletePost(ctx, postID)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}
	utils.Respond(w, http.StatusOK, utils.Message("Post deleted successfully"))
}



