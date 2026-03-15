package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductSizeService struct {
	repo *repository.ProductSizeRepository
}

func NewProductSizeService(rp *repository.ProductSizeRepository) *ProductSizeService {
	return &ProductSizeService{
		repo: rp,
	}
}

func (s *ProductSizeService) Create(req *models.CreateProductSizeRequest) error {
	if req.ProductId < 0 {
		return errors.New("Product ID invalid")
	}
	if req.SizeId < 0 {
		return errors.New("Size ID invalid")
	}

	p := models.ProductSize{
		ProductId: req.ProductId,
		SizeId:    req.SizeId,
	}
	return s.repo.Create(p)
}

func (s *ProductSizeService) GetAll() ([]models.ProductSize, error) {
	return s.repo.GetAll()
}

func (s *ProductSizeService) GetById(id int) (*models.ProductSize, error) {
	return s.repo.GetById(id)
}
