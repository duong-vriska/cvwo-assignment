package main

import (
	"database/sql"
	"net/http"
	"log"
	f "fmt"
	//"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/joho/godotenv"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5"
	

    db "github.com/duong-vriska/cvwo-assignment/backend/database"
	. "github.com/duong-vriska/cvwo-assignment/backend/domain/posts"
	"github.com/duong-vriska/cvwo-assignment/backend/config"
)

type Server struct {
	Version string
	db *db.Queries
	router chi.Router
	httpServer *http.Server
}

func RunServer() *Server{
	s := &Server{
		Version: "1.0.0",
		router: SetupServer(),
	}

	//http.ListenAndServe(":4000", SetupServer())
	f.Println("Server is running on port 4000")
	return s
}

func (s *Server) RunDatabase(){

    database, err := sql.Open(config.DbDriver(), config.DbSource())
	if err != nil {
		panic(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	}

    queries := db.New(database)
    s.db = queries

	f.Printf("Connected to %s database at %s\n", config.DbDriver(), config.DbSource())

    // Perform DB migration
    DBMigration(config.MigrationURL(), config.DbSource())
	defer database.Close()
	
}

func DBMigration(migrationURL string, dbSource string) {
    migration, err := migrate.New(migrationURL, "mysql://" + dbSource)
    if err != nil {
        log.Fatal(err)
    }

    if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal("An error occurred while syncing the database: ", err) 
    }

    f.Println("Migrated!")
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
	s := RunServer()
	s.RunDatabase()
	s.Reroute()
}



