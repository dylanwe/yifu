package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Mode string

const (
	ModeDev  Mode = "dev"
	ModeProd Mode = "prod"
)

type Config struct {
	Mode       Mode
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
}

var config *Config

func Init() {
	godotenv.Load("../.env")
	modeEnv := os.Getenv("MODE")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	var mode Mode

	switch modeEnv {
	case string(ModeProd):
		mode = ModeProd
	default:
		mode = ModeDev
	}

	config = &Config{
		Mode:       mode,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		DBName:     dbName,
	}
}

func GetConfig() *Config {
	return config
}
