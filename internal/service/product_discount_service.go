package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductDiscountService struct {
	repo *repository.ProductDiscountRepository
}

func NewProductDiscountService(rp *repository.ProductDiscountRepository) *ProductDiscountService {
	return &ProductDiscountService{
		repo: rp,
	}
}

func (s *ProductDiscountService) Create(req *models.CreateProductDiscountRequest) error {
	if req.ProductId < 0 {
		return errors.New("Product ID invalid")
	}
	if req.DiscountId < 0 {
		return errors.New("Discount ID invalid")
	}

	p := models.ProductDiscount{
		ProductId:  req.ProductId,
		DiscountId: req.DiscountId,
	}
	return s.repo.Create(p)
}

func (s *ProductDiscountService) GetAll() ([]models.ProductDiscount, error) {
	return s.repo.GetAll()
}

func (s *ProductDiscountService) GetById(id int) (*models.ProductDiscount, error) {
	return s.repo.GetById(id)
}
