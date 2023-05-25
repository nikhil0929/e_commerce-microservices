package dao

import (
	"e_commerce-microservices/DB/postgres"
	"e_commerce-microservices/src/cart/models"
)

type CartDao struct {
	connection *postgres.DB
}

func NewCartDao(connection *postgres.DB) *CartDao {
	return &CartDao{
		connection: connection,
	}
}

// Query returns all carts in the DB with specified query parameters
func (pd *CartDao) Query(queryParams map[string][]string) ([]models.Cart, bool) {
	var RecievedCarts []models.Cart
	rp, err := pd.connection.QueryRecordWithMapConditions(&models.Cart{}, RecievedCarts, queryParams)
	RecievedCarts = rp.([]models.Cart)
	return RecievedCarts, err
}

// Create creates a cart in the DB with specified query parameters and new fields
func (pd *CartDao) CreateCart(newFields models.Cart) bool {
	return pd.connection.CreateRecord(&newFields)
}

// CreateCartItem creates a cart item in the DB with specified query parameters and new fields
func (pd *CartDao) CreateCartItem(newFields models.CartItem) bool {
	return pd.connection.CreateRecord(&newFields)
}

// Update updates a cart in the DB with specified query parameters and new fields
func (pd *CartDao) Update(conditions map[string][]string, newFields models.Cart) bool {
	return pd.connection.UpdateRecord(models.Cart{}, conditions, newFields)
}

// DeleteCart deletes a cart in the DB with specified query parameters
func (pd *CartDao) Delete(conditions map[string][]string) bool {
	return pd.connection.DeleteRecord(models.Cart{}, conditions)
}
