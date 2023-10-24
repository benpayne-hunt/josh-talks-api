package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    error := godotenv.Load()
    if error != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("MONGO_URI")
}