package db

import (
	"fmt"
	"log"
	"os"

	"github.com/elkrammer/gorsvp/internal/config"

	//    "github.com/elkrammer/gorsvp/model"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

func Init() {
	// load config variables from .env file
	config.LoadEnv()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err = sqlx.Connect("pgx", connectionString)
	if err != nil {
		fmt.Println("Failed to connect to database: " + connectionString)
		os.Exit(1)
	}

	if os.Getenv("RUN_MIGRATIONS") == "true" {
		RunMigrations()
	}
}

func RunMigrations() {
	config.LoadEnv()
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	m, err := migrate.New("file://db/migrations", connectionString)
	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("an error occurred while syncing the database.. %v", err)
	}

}

func DbManager() *sqlx.DB {
	if db == nil {
		panic("Failed to get DB Connection")
	}
	return db
}
