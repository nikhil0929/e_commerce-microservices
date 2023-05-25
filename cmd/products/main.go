package main

import (
	"e_commerce-microservices/DB/postgres"
	"e_commerce-microservices/src/products/api"
	"e_commerce-microservices/src/products/dao"
	"e_commerce-microservices/src/products/models"
	"e_commerce-microservices/src/products/router"
	"e_commerce-microservices/src/products/services"
	"e_commerce-microservices/utils"
	"log"
)

func main() {

	utils.LoadEnv("./src/products/.env")

	// Declare Postgres instance
	product_psql := postgres.NewDBConnection_postgres()

	// Migrate products model to a postgres table
	product_psql.RunMigrations(models.Product{})

	// Declare product service
	product_dao := dao.NewProductDao(product_psql)
	product_services := services.NewProductService(product_dao)
	product_controller := api.NewProductController(product_services)

	product_app := router.NewProductService(product_controller)

	log.Println("Starting Product Microservice on port 4001")

	product_app.Start()

}
