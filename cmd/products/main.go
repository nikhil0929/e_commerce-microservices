package main

import (
	"e_commerce-microservices/DB/postgres"
	"e_commerce-microservices/products/api"
	"e_commerce-microservices/products/dao"
	"e_commerce-microservices/products/router"
	"e_commerce-microservices/products/services"
	"log"
)

func main() {

	// Declare product service
	product_psql := postgres.NewDBConnection_postgres()
	product_dao := dao.NewProductDao(product_psql)
	product_services := services.NewProductService(product_dao)
	product_controller := api.NewProductController(product_services)

	product_app := router.NewProductService(product_controller)

	log.Println("Starting Product Microservice on port 4001")

	product_app.Start()

}