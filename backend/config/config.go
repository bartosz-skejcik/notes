package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	POSTGRES_HOST string
	POSTGRES_PORT string
	POSTGRES_USER string
	POSTGRES_PASSWORD string
	POSTGRES_DATABASE string
	POSTGRES_SSL string

	REDIS_HOST string
	REDIS_PASSWORD string
}

func LoadEnvVars() EnvVars {
	var envVars EnvVars
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	envVars.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	envVars.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	envVars.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	envVars.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	envVars.POSTGRES_DATABASE = os.Getenv("POSTGRES_DATABASE")
	envVars.POSTGRES_SSL = os.Getenv("POSTGRES_SSL")

	envVars.REDIS_HOST = os.Getenv("REDIS_HOST")
	envVars.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

	return envVars
}