package postHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"fmt"
	_ "github.com/duong-vriska/cvwo-assignment/backend/posts"
)

type PostHandler struct {
	storage PostStore
}

func (b PostHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(b.storage.List())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	post := b.storage.Get(id)
	if post == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	err := json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b.storage.Create(post)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedPost := b.storage.Update(id, post)
	if updatedPost == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(updatedPost)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	post := b.storage.Delete(id)
	if post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
