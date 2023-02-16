package services

import (
	"e_commerce-microservices/products/models"
	"e_commerce-microservices/utils"
)

type DAO interface {
	Query(map[string][]string) ([]models.Product, bool)
	Create(models.Product) bool
	Update(map[string][]string, models.Product) bool
	Delete(map[string][]string) bool
}

type Service struct {
	productDao DAO
}

func NewProductService(productDao DAO) *Service {
	return &Service{
		productDao: productDao,
	}
}

func (s *Service) GetProducts(queryParams map[string][]string) ([]models.Product, bool) {
	prods, err := s.productDao.Query(queryParams)
	return prods, err
}

func (s *Service) CreateProduct(prod models.Product) bool {
	// make sure product is valid
	isValid := utils.CheckProductValidity(prod)
	if !isValid {
		return false
	}
	return s.productDao.Create(prod)
}

func (s *Service) UpdateProduct(conditions map[string][]string, newFields models.Product) bool {
	// make sure product is valid
	isValid := utils.CheckProductValidity(newFields)
	if !isValid {
		return false
	}
	return s.productDao.Update(conditions, newFields)
}

func (s *Service) DeleteProduct(conditions map[string][]string) bool {
	return s.productDao.Delete(conditions)
}