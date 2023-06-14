package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_URI")
}

func GetEnv(ENV_NAME string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error: Unable to load %s variable.", ENV_NAME)
	}

	return os.Getenv(ENV_NAME)
}
