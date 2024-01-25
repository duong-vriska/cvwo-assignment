package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DbSource() string {
	currentDir, err := os.Getwd()
	envFile, err := godotenv.Read(currentDir + "/.env")
	if err != nil {
		log.Fatal(err)
	}
	
	dbSource := envFile["DB_SOURCE"]

    if dbSource == "" {
        fmt.Println("DB_SOURCE environment variable not set")
    } else {
        fmt.Println("OK SOURCE!")
    }

    return dbSource
}

func DbDriver() string {
    return "mysql"
}

func MigrationURL() string {
	currentDir, err := os.Getwd()
	envFile, err := godotenv.Read(currentDir + "/.env")
	if err != nil {
		log.Fatal(err)
	}
	
	migrationURL := envFile["MIGRATION_URL"]

    if migrationURL == "" {
        fmt.Println("MIGRATION_URL environment variable not set")
    } else {
        fmt.Println("OK URL!")
    }

    return migrationURL

}