package localenv

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func readEnv() {

	var envs map[string]string
	envs, err := godotenv.Read("local.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := envs["DB_NAME"]
	dbUser := envs["DB_USER"]
	dbPass := envs["DB_PASS"]
	dbPort := envs["DB_PORT"]

	fmt.Printf("name %s user %s pw %s port %s\n", dbName, dbUser, dbPass, dbPort)
}
