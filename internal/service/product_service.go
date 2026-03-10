package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(rp *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: rp,
	}
}

func (s *ProductService) CreateProduct(req *models.CreateProductRequest) error {
	if req.Price <= 0 {
		return errors.New("Price must be greater than zero")
	}

	product := models.Product{
		ProductName: req.ProductName,
		Desc:        req.Desc,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	return s.repo.CreateProduct(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) GetProductById(id int) (*models.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *ProductService) Update(id int, req models.UpdateProductRequest) error {
	existing, err := s.repo.GetProductById(id)

	if err != nil {
		return errors.New("product not found")
	}
	if req.ProductName != "" {
		existing.ProductName = req.ProductName
	}
	if req.Desc != "" {
		existing.Desc = req.Desc
	}
	if req.Price > 0 {
		existing.Price = req.Price
	}
	if req.Stock > 0 {
		existing.Stock = req.Stock
	}

	return s.repo.Update(id, *existing)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
