package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	DatabaseType string
	KeycloakURL  string
	KeycloakRealm string
	KeycloakClientID string
	KeycloakClientSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		DatabaseType: os.Getenv("DATABASE_TYPE"),
		KeycloakURL:  os.Getenv("KEYCLOAK_URL"),
		KeycloakRealm: os.Getenv("KEYCLOAK_REALM"),
		KeycloakClientID: os.Getenv("KEYCLOAK_CLIENT_ID"),
		KeycloakClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
	}
}