package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductVariantService struct {
	repo *repository.ProductVariantRepository
}

func NewProductVariantService(rp *repository.ProductVariantRepository) *ProductVariantService {
	return &ProductVariantService{
		repo: rp,
	}
}

func (s *ProductVariantService) Create(req *models.CreateProductVariantRequest) error {
	if req.ProductId < 0 {
		return errors.New("Product ID invalid")
	}
	if req.VariantId < 0 {
		return errors.New("Variant ID invalid")
	}

	p := models.ProductVariant{
		ProductId: req.ProductId,
		VariantId: req.VariantId,
	}
	return s.repo.Create(p)
}

func (s *ProductVariantService) GetAll() ([]models.ProductVariant, error) {
	return s.repo.GetAll()
}

func (s *ProductVariantService) GetById(id int) (*models.ProductVariant, error) {
	return s.repo.GetById(id)
}
