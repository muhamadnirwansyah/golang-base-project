package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   Server
	Database Database
	Secret   Secret
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Name    string
	Tz      string
	Migrate string
}

type Secret struct {
	Jwt string
}

func Get() *Config {
	fileFlag := flag.String("env", "", "file .env location path absolute")
	flag.Parse()

	var err error
	if *fileFlag != "" {
		err = godotenv.Load(*fileFlag)
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal("Error when load .env : ", err.Error())
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host:    os.Getenv("DB_HOST"),
			Port:    os.Getenv("DB_PORT"),
			User:    os.Getenv("DB_USER"),
			Pass:    os.Getenv("DB_PASS"),
			Name:    os.Getenv("DB_NAME"),
			Tz:      os.Getenv("DB_TIMEZONE"),
			Migrate: os.Getenv("MIGRATE_DATABASE_FROM_DOMAIN"),
		},
		Secret: Secret{
			Jwt: os.Getenv("SECRET_JWT"),
		},
	}
}
