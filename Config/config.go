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

var Host string
var Db_username string
var Db_password string
var Port int
var Db_name string

var JWT_secret_key string
var Cookie_secret_key string
var Stripe_secret_key string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	Host = os.Getenv("HOST")
	Db_username = os.Getenv("DB_USERNAME")
	Db_password = os.Getenv("DB_PASSWORD")
	Port, _ = strconv.Atoi(os.Getenv("PORT"))
	Db_name = os.Getenv("DB_NAME")

	JWT_secret_key = os.Getenv("JWT_SECRET_KEY")
	Cookie_secret_key = os.Getenv("COOKIE_SECRET_KEY")
	Stripe_secret_key = os.Getenv("STRIPE_SECRET_KEY")

}