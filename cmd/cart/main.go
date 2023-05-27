package main

import (
	"e_commerce-microservices/src/cart/dao"
	"e_commerce-microservices/src/cart/models"
	"e_commerce-microservices/utils"
	"fmt"
)

func main() {

	utils.LoadEnv("./src/cart/.env")

	// Declare Cart service
	cart_dao := dao.NewCartDao()

	// Migrate Cart model to a postgres table
	cart_dao.RunMigrations(models.CartItem{})

	// Migrate Cart model to a postgres table
	cart_dao.RunMigrations(models.Cart{})

	// cart_services := services.NewCartService(cart_dao)

	// prod, err := cart_services.GetProduct(1)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // Print the product
	// fmt.Printf("%+v\n", prod)

	// isValid := cart_services.InsertCartItem(1, 1, 2)

	// fmt.Println(isValid)

	// CartItem represents a single item in the cart
	// type CartItem struct {
	// 	ID         uint `gorm:"primary_key"`
	// 	CartID     uint
	// 	ProductID  uint
	// 	TotalPrice float64
	// 	Quantity   int
	// }

	// cart_item := &models.CartItem{
	// 	CartID:     1,
	// 	ProductID:  1,
	// 	TotalPrice: 100,
	// 	Quantity:   2,
	// }

	// cart_dao.CreateRecord(cart_item)

	// query cart item
	cart_items := &[]models.CartItem{}
	cart_dao.QueryRecords(cart_items, map[string]interface{}{"cart_id": 1})

	// Print the cart_items
	fmt.Printf("%+v\n", cart_items)

	first_cart_item := (*cart_items)[0]

	cart_dao.UpdateRecord(first_cart_item, map[string]interface{}{"quantity": 3})

	cart_dao.QueryRecords(cart_items, map[string]interface{}{"cart_id": 1})

	// Print the cart_items
	fmt.Printf("%+v\n", cart_items)

	cart_dao.DeleteRecord(first_cart_item, map[string]interface{}{})

}
