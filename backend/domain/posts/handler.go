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
	return &Handler{db: db}
}

func (h *Handler) HandleListPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := h.db.ListPost(ctx)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, "Can't get")
		return
	}
	utils.WriteJSON(w, http.StatusOK, posts)
}

func (h *Handler) HandleGetPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := chi.URLParam(r, "post_id")
	post, err := h.db.GetPost(ctx, postID)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, "Can't get")
		return
	}
	utils.WriteJSON(w, http.StatusOK, post)
}

func (h *Handler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var post db.CreatePostParams
	utils.Parse(w, r, &post)
	post.PostID = utils.GenerateRandomString(5)
	createdPost, err := h.db.CreatePost(ctx, post)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, "Can't create")
		return
	}
	utils.WriteJSON(w, http.StatusOK, createdPost)
}

func (h *Handler) HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var post db.UpdatePostParams
	utils.Parse(w, r, &post)
	updatedPost, err := h.db.UpdatePost(ctx, post)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, "Can't update")
		return
	}
	utils.WriteJSON(w, http.StatusOK, updatedPost)
}

func (h *Handler) HandleDeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := chi.URLParam(r, "post_id")
	err := h.db.DeletePost(ctx, postID)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, "Can't delete")
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "deleted"})
}


