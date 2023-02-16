package products

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetProducts(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}


type Product_Mircoservice struct {
	api Controller
}

func NewProductService(api Controller) *Product_Mircoservice {
	return &Product_Mircoservice{
		api: api,
	}
}

func enableProductRoutes(router *gin.Engine, ps *Product_Mircoservice) {

	// Product routes
	router.GET("/products", ps.api.GetProducts)

	// all routes here are for admin user group only
	router.POST("/products", ps.api.CreateProduct)
	router.PUT("/products", ps.api.UpdateProduct)
	router.DELETE("/products", ps.api.DeleteProduct)
}

func (ps *Product_Mircoservice) newProductsRouter() *gin.Engine {
	router := gin.Default()

	enableProductRoutes(router, ps)
	
	return router
}

func (ps *Product_Mircoservice) Start() {
	router := ps.newProductsRouter()
	log.Println("Starting Product Microservice on port 4001")
	router.Run(":4001")
}