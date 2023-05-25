package main

import (
	"e_commerce-microservices/DB/postgres"
	"e_commerce-microservices/src/users/api"
	"e_commerce-microservices/src/users/dao"
	"e_commerce-microservices/src/users/models"
	"e_commerce-microservices/src/users/router"
	"e_commerce-microservices/src/users/services"
	"e_commerce-microservices/utils"
	"log"
)

func main() {

	utils.LoadEnv("./src/users/.env")

	// Declare Postgres instance
	user_psql := postgres.NewDBConnection_postgres()

	// Migrate users model to a postgres table
	user_psql.RunMigrations(models.User{})

	// Declare user service
	user_dao := dao.NewUserDao(user_psql)
	user_services := services.NewUserService(user_dao)
	user_controller := api.NewUserController(user_services)

	user_app := router.NewUserService(user_controller)

	log.Println("Starting User Microservice on port 4002")

	user_app.Start()

}
