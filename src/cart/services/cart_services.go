package services

import (
	"e_commerce-microservices/src/cart/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type DAO interface {
	RunMigrations(interface{})
	QueryRecords(interface{}, map[string]interface{}) error
	QueryWithAssociation(interface{}, map[string]interface{}, string) error
	CreateRecord(interface{}) error
	UpdateRecord(interface{}, map[string]interface{}) error
	DeleteRecord(interface{}, map[string]interface{}) error
}

type Service struct {
	cartDao DAO
}

func NewCartService(dao DAO) *Service {
	return &Service{
		cartDao: dao,
	}
}

// Cart Service methods

func (s *Service) GetCart(queryParams map[string][]string) ([]models.Cart, error) {

}

// Creates an empty cart for the user with the users first item. This method should not really be called by the client. It should be called by the 'InsertCartItem' method
func (s *Service) CreateCart(UserID uint, cartItem models.CartItem) error {

}

// This function calls the 'GetProduct' function from the product microservice to get the product object using the product ID

// if the cart for the user does not exist, it will create a new cart for the user
// if the cart for the user exists, it will check if the product already exists in the cart
// if the product exists in the cart, it will update the quantity of the product in the cart
// if the product does not exist in the cart, it will add the product to the cart by first creating a new cart item and then adding it to the cart
func (s *Service) InsertCartItem(ProductID uint, UserID uint, Quantity uint) error {

}

// This function calls the 'GetProduct' function from the product microservice to get the product object using the product ID and constructs a cart item object
func CreateCartItem(ProductID uint, UserID uint, Quantity uint) (models.CartItem, error) {

}

// Deletes an CartItem from the cart
func (s *Service) DeleteCartItem(CartID uint, CartItemID uint) bool {
	return true
}

// Updates an CartItem in the cart
func (s *Service) UpdateCartItem(CartItemID uint) bool {
	return true
}

// calls the 'GetProduct' function from the product microservice to get the product object using the product ID
func GetProduct(productID uint) (*[]models.Product, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:4001/products?ID=%s", strconv.Itoa(int(productID))))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product api returned error: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, fmt.Errorf("could not parse product response: %w", err)
	}

	return &products, nil
}
