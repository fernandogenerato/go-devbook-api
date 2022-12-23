package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	UriConection = ""
	Port         = 0
	SecretKey    []byte
)

type Config struct {
	UriConection string
	Port         int
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	UriConection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	return Config{
		UriConection: UriConection,
		Port:         Port,
	}

}
