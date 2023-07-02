package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(ENV_NAME string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error: Unable to load %s variable.", ENV_NAME)
	}
	envVar := os.Getenv(ENV_NAME)
	if envVar == "" {
		fmt.Println("error: field is empty", ENV_NAME)
	}
	return envVar
}
