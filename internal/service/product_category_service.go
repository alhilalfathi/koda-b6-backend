package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductCategoryService struct {
	repo *repository.ProductCategoryRepository
}

func NewProductCategoryService(rp *repository.ProductCategoryRepository) *ProductCategoryService {
	return &ProductCategoryService{
		repo: rp,
	}
}

func (s *ProductCategoryService) Create(req *models.CreateProductCategoryRequest) error {
	if req.ProductId < 0 {
		return errors.New("Product ID invalid")
	}
	if req.CategoryId < 0 {
		return errors.New("Category ID invalid")
	}

	p := models.ProductCategory{
		ProductId:  req.ProductId,
		CategoryId: req.CategoryId,
	}
	return s.repo.Create(p)
}

func (s *ProductCategoryService) GetAll() ([]models.ProductCategory, error) {
	return s.repo.GetAll()
}

func (s *ProductCategoryService) GetById(id int) (*models.ProductCategory, error) {
	return s.repo.GetById(id)
}
