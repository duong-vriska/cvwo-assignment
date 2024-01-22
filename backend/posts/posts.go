package posts

import (
	db "github.com/duong-vriska/cvwo-assignment/backend/database"
	. "github.com/duong-vriska/cvwo-assignment/backend/utils"
)

type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

type PostStorage interface {
	List() []*Post
	Get(string) *Post
	Update(string, Post) *Post
	Create(Post)
	Delete(string) *Post
}

type PostStore struct{}

var posts = []*Post{
	{
		ID:      "1",
		Title:   "The first post",
		Content: "Mowou is a great place to discuss and share your ideas",
		Category: "Announcement",
	},
}

func (b PostStore) Get(id string) *Post {
	for _, post := range posts {
		if post.ID == id {
			return post
		}
	}
	return nil
}

func (b PostStore) List() []*Post {
	return posts
}

func (b PostStore) Create(post Post) {
	post.ID = GenerateRandomString(5) 
	posts = append(posts, &post)
}

func (b PostStore) Delete(id string) *Post {
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			return &Post{}
		}
	}
	return nil
}

func (b PostStore) Update(id string, postUpdate Post) *Post {
	for i, post := range posts {
		if post.ID == id {
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
		if post.ID == id {
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
		if post.ID == id {
			posts = append(posts[:i], (posts)[i+1:]...)
			return &Post{}
		}
	}
	return nil
}
func updatePost(id string, postUpdate Post) *Post {
	for i, post := range posts {
		if post.ID == id {
			posts[i] = &postUpdate
			return post
		}
	}
	return nil
}
