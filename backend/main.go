package main

import (
    "context"
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
    db "github.com/duong-vriska/cvwo-assignment/backend/database/schema"
)

func run() error {
    ctx := context.Background()

    database, err := sql.Open("mysql", "root:nmai1202@/db?parseTime=true")
    if err != nil {
        return err
    }

    queries := db.New(database)

    // list all authors
    posts, err := queries.ListPost(ctx)
    if err != nil {
        return err
    }
    log.Println(posts)

    // prints true
    //log.Println(reflect.DeepEqual(insertedAuthorID, fetchedAuthor.ID))
    return nil
}

func main() {
    if err := run(); err != nil {
        log.Fatal(err)
    }
}
