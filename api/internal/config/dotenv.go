package config

import (
    "fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
    _, err := os.Stat(".env")
    if !os.IsNotExist(err) {
        err := godotenv.Load()
        if err != nil {
            err = godotenv.Load(".env")
            if err != nil {
                log.Fatal("Error loading .env file")
            }
        }
    }

    fmt.Println("========== ENV VARS ==========")
	vars := []string{"DB_USERNAME", "DB_HOST", "DB_PASSWORD", "DB_NAME", "DB_PORT", "JWT_SECRET", "SENDGRID_API_KEY", "RUN_MIGRATIONS"}
	for _, k := range vars {
        value := os.Getenv(k)
        fmt.Printf("%s: %s\n", k, value)
    }
    fmt.Println("========== ENV VARS ==========")

	if vars := checkVars(); len(vars) != 0 {
		log.Printf("Error: Mandatory environment variables not set in .env file: %v", vars)
	}
}

func checkVars() []string {
	vars := []string{"DB_USERNAME", "DB_PASSWORD", "DB_HOST", "DB_NAME", "DB_PORT", "JWT_SECRET", "SENDGRID_API_KEY", "RUN_MIGRATIONS"}
	missing := []string{}
	for _, v := range vars {
		_, set := os.LookupEnv(v)
		if !set {
			missing = append(missing, v)
		}
	}
	return missing
}
