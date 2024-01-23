package main

import (
    "context"
	"database/sql"
	"net/http"
    "log"
	f "fmt"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5"
    _ "github.com/go-sql-driver/mysql"
    db "github.com/duong-vriska/cvwo-assignment/backend/database"
	. "github.com/duong-vriska/cvwo-assignment/backend/domain/posts"
)

type Server struct {
	Version string

	db *db.Queries

	router chi.Router
	httpServer *http.Server
}

func RunServer(version string, db *db.Queries) *Server {
	s := &Server{
		Version: version,
		db: db,
	}

	//http.ListenAndServe(":4000", SetupServer())
	f.Println("Server is running on port 4000")
	return s
}

func (s *Server) RunDatabase(){
	ctx := context.Background()

	database, err := sql.Open("mysql", "root:nmai1202@(127.0.0.1:3306)/db?parseTime=true")
	if err != nil {
		f.Println("Failed to connect to database!")
	}

	err = database.Ping()
	if err != nil {
		f.Println("Failed to ping database!")
	}

	queries := db.New(database)
	s.db = queries 

	posts, err := queries.ListPost(ctx)
	if err != nil {
		return 
	}
	log.Println(posts)

	return 
}

func SetupServer() chi.Router {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(cors.Handler(cors.Options{
        // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
        AllowedOrigins:   []string{"https://*", "http://*"},
        // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
      }))

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        _, err := w.Write([]byte("OK"))
        if err != nil {
            return
        }
    })
    return r
}

func (s *Server) Reroute() {
	s.router.Mount("/posts", RunRouter(s.db))
}

func main(){
	s := RunServer("1.0.0", nil)
	s.RunDatabase()
	//s.Reroute()
}



