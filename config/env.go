package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// "LocalEnv - load local environment variables"
func LocalEnv() (string, string, string) {

	var envs map[string]string
	envs, err := godotenv.Read("local.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := envs["POSTGRES_DB"]
	dbUser := envs["POSTGRES_USER"]
	dbPass := envs["POSTGRES_PASSWORD"]

	fmt.Printf("name %s user %s pw %s", dbName, dbUser, dbPass)
	return dbName, dbUser, dbPass
}
