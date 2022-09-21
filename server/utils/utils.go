package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
	return os.Getenv(key)
}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

type HttpResp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
