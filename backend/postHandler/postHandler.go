package postHandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	// . "github.com/duong-vriska/cvwo-assignment/backend/posts"
	//"github.com/duong-vriska/cvwo-assignment/backend/database/schema"
	"github.com/duong-vriska/cvwo-assignment/backend/utils"
)

type Post struct {
	ID       int32  `json:"id"`
	PostID   string `json:"post_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

var posts = []*Post{
    {
        ID:      1,
        PostID:  "MnsSh",
        Title:   "The first post",
        Content: "Mowou is a great place to discuss and share your ideas",
        Category: "Announcement",
    },
}

type PostStore struct{}

type PostStorage interface {
    List() []*Post
    Get(string) *Post
    Update(string, Post) *Post
    Create(Post)
    Delete(string) *Post
}

type PostHandler struct {
	storage PostStorage
}

func (b PostStore) Get(id string) *Post {
    for _, post := range posts {
        if post.PostID == id {
            return post
        }
    }
    return nil
}

func (b PostStore) List() []*Post {
    return posts
}

func (b PostStore) Create(post Post) {
    post.PostID = utils.GenerateRandomString(5) 
    posts = append(posts, &post)
}

func (b PostStore) Delete(id string) *Post {
    for i, post := range posts {
        if post.PostID == id {
            posts = append(posts[:i], posts[i+1:]...)
            return &Post{}
        }
    }
    return nil
}

func (b PostStore) Update(id string, postUpdate Post) *Post {
    for i, post := range posts {
        if post.PostID == id {
            posts[i] = &postUpdate
            return post
        }
    }
    return nil
}

func listPosts() []*Post {
    return posts
}

func getPost(id string) *Post {
    for _, post := range posts {
        if post.PostID == id {
            return post
        }
    }
    return nil
}

func storePost(post Post) {
    posts = append(posts, &post)
}

func deletePost(id string) *Post {
    for i, post := range posts {
        if post.PostID == id {
            posts = append(posts[:i], (posts)[i+1:]...)
            return &Post{}
        }
    }
    return nil
}
func updatePost(id string, postUpdate Post) *Post {
    for i, post := range posts {
        if post.PostID == id {
            posts[i] = &postUpdate
            return post
        }
    }
    return nil
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

	posts := []*Post{post}

	err := json.NewEncoder(w).Encode(posts)
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
    if id == "" {
        http.Error(w, "Invalid or missing 'id' parameter", http.StatusBadRequest)
        return
    }
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

func (b PostHandler) FilterPost(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	posts := b.storage.List()
	filteredPosts := []*Post{}
	for _, post := range posts {
		if post.Category == category {
			filteredPosts = append(filteredPosts, post)
		}
	}
	err := json.NewEncoder(w).Encode(filteredPosts)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}