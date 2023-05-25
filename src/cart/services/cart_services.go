package services

import (
	"e_commerce-microservices/src/cart/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type DAO interface {
	Query(map[string][]string) ([]models.Cart, bool)
	CreateCart(models.Cart) bool
	CreateCartItem(models.CartItem) bool
	Update(map[string][]string, models.Cart) bool
	Delete(map[string][]string) bool
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

func (s *Service) GetCart(queryParams map[string][]string) ([]models.Cart, bool) {
	carts, err := s.cartDao.Query(queryParams)
	return carts, err
}

// Creates an empty cart for the user with the users first item. This method should not really be called by the client. It should be called by the 'InsertCartItem' method
func (s *Service) CreateCart(UserID uint, cartItem models.CartItem) bool {
	cart := models.Cart{
		UserID: UserID,
		CartItems: []models.CartItem{
			cartItem,
		},
	}
	isValid := s.cartDao.CreateCart(cart)
	return isValid
}

// This function calls the 'GetProduct' function from the product microservice to get the product object using the product ID

// if the cart for the user does not exist, it will create a new cart for the user
// if the cart for the user exists, it will check if the product already exists in the cart
// if the product exists in the cart, it will update the quantity of the product in the cart
// if the product does not exist in the cart, it will add the product to the cart by first creating a new cart item and then adding it to the cart
func (s *Service) InsertCartItem(ProductID uint, UserID uint, Quantity uint) bool {
	// Check if cart exists for user
	carts, err := s.cartDao.Query(map[string][]string{"User_ID": {strconv.Itoa(int(UserID))}})

	if err == false {
		log.Println("InsertCartItem: Error querying cart - ", err)
		return false
	}

	cartItem, isValid := CreateCartItem(ProductID, UserID, Quantity)

	if isValid == false {
		log.Println("InsertCartItem: Error creating cart item - ", err)
		return false
	}

	if len(carts) == 0 {
		// Create cart for user

		isValid = s.CreateCart(UserID, cartItem)

		if isValid == false {
			log.Println("InsertCartItem: Error creating cart - ", err)
			return false
		}

		return true
	} else {
		// Check if product already exists in cart

		cart := carts[0]

		for _, item := range cart.CartItems {
			if item.ProductID == ProductID {
				// Update quantity of product in cart

				isValid = s.UpdateCartItem(item.ID)

				if isValid == false {
					log.Println("InsertCartItem: Error updating cart item - ", err)
					return false
				}

				return true
			}
		}

		// Create a new cart item and add it to the cart

		isValid = s.cartDao.CreateCartItem(cartItem)

		if isValid == false {
			log.Println("InsertCartItem: Error creating cart item - ", err)
			return false
		}

		return true

	}
}

// This function calls the 'GetProduct' function from the product microservice to get the product object using the product ID and constructs a cart item object
func CreateCartItem(ProductID uint, UserID uint, Quantity uint) (models.CartItem, bool) {
	products, err := GetProduct(ProductID)

	if err != nil {
		log.Println("CreateCartItem: Error getting product - ", err)
		return models.CartItem{}, false
	}

	product := (*products)[0]

	// type CartItem struct {
	// 	CartID    uint `gorm:"primary_key"`
	// 	ProductID uint
	// 	TotalPrice     float64
	// 	Quantity  int
	// }

	totalPrice := product.Price * float64(Quantity)
	cartItem := models.CartItem{
		ProductID:  product.ID,
		TotalPrice: totalPrice,
		Quantity:   int(Quantity),
	}

	return cartItem, true
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
