package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/elkrammer/kegger/api/internal/config"
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

	// configure sql.DB for better performance
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

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
