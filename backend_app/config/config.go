package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST        string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_NAME        string
	DB_PORT        string
	SERVER_ADDRESS string
}

func init() {
	err := godotenv.Load()
	if(err != nil){
		log.Fatal("Error load .env file")
	}
}

func InitConfiguration() Config{
	return Config{
		DB_HOST: os.Getenv("DB_HOST"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME: os.Getenv("DB_NAME"),
		DB_PORT: os.Getenv("DB_PORT"),
		SERVER_ADDRESS: os.Getenv("SERVER_ADDRESS"),
	}
}