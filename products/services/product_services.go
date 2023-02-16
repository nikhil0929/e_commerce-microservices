package services

import (
	"e_commerce-microservices/products/models"
	"e_commerce-microservices/utils"
)

type DAO interface {
	Query(map[string][]string) []models.Product
	Create(*models.Product) bool
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

func (s *Service) GetProducts(queryParams map[string][]string) []models.Product {
	prods := s.productDao.Query(queryParams)
	return prods
}

func (s *Service) CreateProduct(prod *models.Product) {
	// make sure product is valid
	// make sure product doesnt exist already
	isValid := utils.CheckProductValidity(*prod)
	if !isValid {
		return
	s.productDao.Create(prod)
	return
}

func (s *Service) UpdateProduct(conditions map[string][]string, newFields models.Product) bool {
	// make sure product is valid
	// make sure product doesnt exist already
	s.productDao.Update(prod)
	return
}

func (s *Service) DeleteProduct(conditions map[string][]string) {
	// make sure product is valid
	// make sure product doesnt exist already
	s.productDao.Delete(prod)
	return
}