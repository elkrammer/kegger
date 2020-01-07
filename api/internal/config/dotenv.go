package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	if vars := checkVars(); len(vars) != 0 {
		fmt.Sprintf("Error: Mandatory environment variables not set in .env file: %v", vars)
	}
}

func checkVars() []string {
	vars := []string{"DB_USERNAME", "DB_PASSWORD", "DB_NAME", "DB_PORT", "JWT_SECRET"}
	missing := []string{}
	for _, v := range vars {
		_, set := os.LookupEnv(v)
		if !set {
			missing = append(missing, v)
		}
	}
	return missing
}
