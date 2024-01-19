package posts

type Post struct {
	ID      int `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostStorage interface {
	List() []*Post
	Get(int) *Post
	Update(int, Post) *Post
	Create(Post)
	Delete(int) *Post
}

type PostStore struct{}

var posts = []*Post{
	{
		ID:      1,
		Title:   "Delicious Chicken",
		Content: "Lorem ipsum",
	},
}

func (b PostStore) Get(id int) *Post {
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
	posts = append(posts, &post)
}

func (b PostStore) Delete(id int) *Post {
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			return &Post{}
		}
	}
	return nil
}

func (b PostStore) Update(id int, postUpdate Post) *Post {
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

func getPost(id int) *Post {
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

func deletePost(id int) *Post {
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], (posts)[i+1:]...)
			return &Post{}
		}
	}
	return nil
}
func updatePost(id int, postUpdate Post) *Post {
	for i, post := range posts {
		if post.ID == id {
			posts[i] = &postUpdate
			return post
		}
	}
	return nil
}
