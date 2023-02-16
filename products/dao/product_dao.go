package dao

import (
	"e_commerce-microservices/DB/postgres"
	"e_commerce-microservices/products/models"
)

type ProductDao struct{
	connection *postgres.DB
}

func NewProductDao(connection *postgres.DB) *ProductDao {
	return &ProductDao{
		connection: connection,
	}
}

// Query returns all products in the DB with specified query parameters
func (pd *ProductDao) Query(queryParams map[string][]string) ([]models.Product, bool) {
	var RecievedProducts []models.Product
	rp, err := pd.connection.QueryRecordWithMapConditions(&models.Product{}, RecievedProducts, queryParams)
	RecievedProducts = rp.([]models.Product)
	return RecievedProducts, err
}

// Create creates a product in the DB with specified query parameters and new fields
func (pd *ProductDao) Create(newFields models.Product) bool {
	return pd.connection.CreateRecord(&newFields)
}

// Update updates a product in the DB with specified query parameters and new fields
func (pd *ProductDao) Update(conditions map[string][]string, newFields models.Product) bool {
	return pd.connection.UpdateRecord(models.Product{}, conditions, newFields)
}

// DeleteProduct deletes a product in the DB with specified query parameters
func (pd *ProductDao) Delete(conditions map[string][]string) bool {
	return pd.connection.DeleteRecord(models.Product{}, conditions)
}