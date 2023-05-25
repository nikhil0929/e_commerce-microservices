package main

import (
	"e_commerce-microservices/DB/postgres"
	"e_commerce-microservices/src/cart/dao"
	"e_commerce-microservices/src/cart/models"
	"e_commerce-microservices/src/cart/services"
	"e_commerce-microservices/utils"
	"fmt"
)

func main() {

	utils.LoadEnv("./src/cart/.env")

	cart_psql := postgres.NewDBConnection_postgres()

	// Migrate Cart model to a postgres table
	cart_psql.RunMigrations(models.CartItem{})

	// Migrate Cart model to a postgres table
	cart_psql.RunMigrations(models.Cart{})

	// Declare Cart service
	cart_dao := dao.NewCartDao(cart_psql)
	cart_services := services.NewCartService(cart_dao)

	// prod, err := cart_services.GetProduct(1)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // Print the product
	// fmt.Printf("%+v\n", prod)

	isValid := cart_services.InsertCartItem(1, 1, 2)

	fmt.Println(isValid)

}
