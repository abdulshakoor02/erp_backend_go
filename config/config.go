package config

import (
	"os"
	"strconv"

	"github.com/abdul/erp_backend/logger"
	"github.com/joho/godotenv"
)

var PORT int
var POSTGRES_HOST string
var DB_NAME string
var POSTGRES_USER string
var POSTGRES_PASSWORD string
var POSTGRES_PORT string
var SECRET_KEY string

func LoadEnv() {
	log := logger.Logger
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to load env")
	}
	PORT, _ = strconv.Atoi(os.Getenv("PORT"))
	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	SECRET_KEY = os.Getenv("SECRET_KEY")
}
