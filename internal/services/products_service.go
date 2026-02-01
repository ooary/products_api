package services

import (
	"products_api/internal/models"
	"products_api/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllService() ([]models.Products, error) {
	return s.repo.GetAll()

}

func (s *ProductService) CreateService(data *models.Products) error {
	return s.repo.CreateProduct(data)
}

func (s *ProductService) GetProductByDetailService(id int) (*models.Products, error) {
	return s.repo.GetProductDetail(id)
}

func (s *ProductService) UpdateProduct(p *models.Products) error {
	return s.repo.UpdateProduct(p)

}

func (s *ProductService) DeleteProductService(id int) bool {
	return s.repo.DeleteProduct(id)
}
