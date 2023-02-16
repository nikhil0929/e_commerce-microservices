package Products

import (
	"e_commerce-microservices/DB/postgres"
	"e_commerce-microservices/products/models"
	"e_commerce-microservices/utils"
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
	rp, err_exists := pd.connection.QueryRecordWithMapConditions(&models.Product{}, RecievedProducts, queryParams)
	RecievedProducts = rp.([]models.Product)
	return RecievedProducts, err_exists
}

// Create creates a product in the DB with specified query parameters and new fields
func (pd *ProductDao) Create(newFields models.Product) bool {
	
	isValid := utils.CheckProductValidity(newFields)
	if isValid {
		return pd.connection.CreateRecord(&newFields)
	}
	return false
}

// Update updates a product in the DB with specified query parameters and new fields
func (pd *ProductDao) Update(conditions map[string][]string, newFields models.Product) bool {
	pd.connection.UpdateRecord(models.Product{}, conditions, newFields)
	return true

}

// DeleteProduct deletes a product in the DB with specified query parameters
func (pd *ProductDao) Delete(conditions map[string][]string) bool {
	pd.connection.DeleteRecord(models.Product{}, conditions)
}