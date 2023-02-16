package products

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetProducts(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}


type ProductService struct {
	api Controller
}

func NewProductService(api Controller) *ProductService {
	return &ProductService{
		api: api,
	}
}

func enableProductRoutes(router *gin.Engine, ps *ProductService) {

	// Product routes
	router.GET("/products", ps.api.GetProducts)

	// all routes here are for admin user group only
	router.POST("/products", ps.api.CreateProduct)
	router.PUT("/products", ps.api.UpdateProduct)
	router.DELETE("/products", ps.api.DeleteProduct)
}

func (ps *ProductService) NewProductsRouter() *gin.Engine {
	router := gin.Default()

	enableProductRoutes(router, ps)
	
	return router
}

func (ps *ProductService) start() {
	router := ps.NewProductsRouter()
	router.Run()
}