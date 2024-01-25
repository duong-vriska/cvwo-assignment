package main

import (
	"database/sql"
	"net/http"
	"log"
	f "fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5/middleware"
	

    db "github.com/duong-vriska/cvwo-assignment/backend/database"
	posts "github.com/duong-vriska/cvwo-assignment/backend/domain/posts"
	"github.com/duong-vriska/cvwo-assignment/backend/config"
)

type Server struct {
	db *db.Queries
}

var s Server

func main(){
	s.RunDatabase()
	r:= chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	pHandler := posts.NewHandler(s.db)

	r.Route("/", func(r chi.Router) {
		r.Mount("/posts", posts.PostRouter(pHandler))
	})

	f.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)
	
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

    DBMigration(config.MigrationURL(), config.DbSource())

	s.db = db.New(database)

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