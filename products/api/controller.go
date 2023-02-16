package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"e_commerce-microservices/products/models"
)

type Service interface {
	GetProducts(map[string][]string) ([]models.Product, bool)
	CreateProduct(models.Product) bool
	UpdateProduct(map[string][]string, models.Product) bool
	DeleteProduct(map[string][]string) bool
}

type api struct {
	product_service Service
}

func NewProductController(product_service Service) *api {
	return &api{
		product_service: product_service,
	}
}

// Public level commands

// Get Products from DB with specified query parameter conditions
// e.g /products?category=electronics
// e.g /products?price=100.99&category=applicances
func (a *api) GetProducts(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	result, isValid := a.product_service.GetProducts(queryParams)
	if !isValid {
		c.String(http.StatusBadRequest, "GetProducts: Unable to get products from DB")
		return
	}
	c.JSON(http.StatusOK, result)
}

// Admin level commands

// Create Product in DB with specified fields in the request body as JSON object
// Product fields are specified in the Models.Product struct
func (a *api) CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.String(http.StatusBadRequest, "CreateProduct: Unable to bind JSON body to object")
		return
	}
	// TODO: Put this in a separate function
	isValid := a.product_service.CreateProduct(newProduct)
	if isValid {
		c.String(http.StatusOK, "CreateProduct: SUCCESS")
	} else {
		c.String(http.StatusBadRequest, "CreateProduct: Invalid Product fields specified")
	}
}

// Update Product in DB with new fields in the request body as JSON object and specified conditions in the query parameters
// e.g. /products?id=1&price=100.99
// JSON Body { "price": 50.00 }
func (a *api) UpdateProduct(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	var newFields models.Product
	if err := c.ShouldBindJSON(&newFields); err != nil {
		c.String(http.StatusBadRequest, "UpdateProduct: Unable to bind JSON body to object")
		return
	}
	// TODO: Put this in a separate function
	isValid := a.product_service.UpdateProduct(queryParams, newFields)
	if isValid {
		c.String(http.StatusOK, "UpdateProduct: SUCCESS")
	} else {
		c.String(http.StatusBadRequest, "UpdateProduct: Invalid Product fields specified")
	}
}

// Delete Product in DB with specified conditions in the query parameters (in the URL)
func (a *api) DeleteProduct(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	isValid := a.product_service.DeleteProduct(queryParams)
	if isValid {
		c.String(http.StatusOK, "DeleteProduct: SUCCESS")
	} else {
		c.String(http.StatusBadRequest, "DeleteProduct: Invalid Product fields specified")
	}
}