package main

import (
	"encoding/json"
	_ "encoding/json"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	_ "net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"testing"
	_ "testing"
)

var fakePosts = []*Post{{
	ID:      "1",
	Title:   "How to be a good-tier student",
	Content: "Study",
}}

type fakeStorage struct {
}

func (s fakeStorage) Get(_ string) *Post {
	return fakePosts[0]
}

func (s fakeStorage) Delete(_ string) *Post {
	return nil
}

func (s fakeStorage) List() []*Post {
	return fakePosts
}

func (s fakeStorage) Create(_ Post) {
	return
}

func (s fakeStorage) Update(string, Post) *Post {
	return fakePosts[1]
}

func TestGetBooksHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	w := httptest.NewRecorder()
	postHandler := PostHandler{
		storage: fakeStorage{},
	}
	postHandler.GetPosts(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	post := Post{}
	json.Unmarshal(data, &post)
	if post.Title != "7 habits of highly effective people" {
		t.Errorf("expected ABC got %v", string(data))
	}
}
