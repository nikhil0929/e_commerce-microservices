// code to load data from the env file
// and from here exports those .env data to be used across the app

// var password string = 'code to load PASSWORD from .env file
// and then to use it, just do Config.password
package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env interface {
	loadConfig()
}

type DBConfig struct {
	host string
	db_username string
	db_password string
	port int
	db_name string
}

type JWTConfig struct {
	jwt_secret_key string
	cookie_secret_key string
}

type StripeConfig struct {
	stripe_secret_key string
}



func (db_config *DBConfig) loadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	db_config.host = os.Getenv("HOST")
	db_config.db_username = os.Getenv("DB_USERNAME")
	db_config.db_password = os.Getenv("DB_PASSWORD")
	db_config.port, _ = strconv.Atoi(os.Getenv("PORT"))
	db_config.db_name = os.Getenv("DB_NAME")

}

func (jwt_config *JWTConfig) loadConfig() {
	jwt_config.jwt_secret_key = os.Getenv("JWT_SECRET_KEY")
	jwt_config.cookie_secret_key = os.Getenv("COOKIE_SECRET_KEY")
}

func (stripe_config *StripeConfig) loadConfig() {
	stripe_config.stripe_secret_key = os.Getenv("STRIPE_SECRET_KEY")
}

