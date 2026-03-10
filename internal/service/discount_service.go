package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type DiscountService struct {
	repo *repository.DiscountRepository
}

func NewDiscountService(rp *repository.DiscountRepository) *DiscountService {
	return &DiscountService{
		repo: rp,
	}
}

func (s *DiscountService) CreateDiscount(req *models.CreateDiscountRequest) error {
	if req.Rate < 0 {
		return errors.New("Input invalid")
	}

	disc := models.Discount{
		Rate:        req.Rate,
		Desc:        req.Desc,
		IsFlashSale: *req.IsFlashSale,
	}
	return s.repo.CreateDiscount(disc)
}

func (s *DiscountService) GetAllDiscounts() ([]models.Discount, error) {
	return s.repo.GetAllDiscounts()
}

func (s *DiscountService) GetDiscountById(id int) (*models.Discount, error) {
	return s.repo.GetDiscountById(id)
}

func (s *DiscountService) Update(id int, req models.UpdateDiscountRequest) error {
	existing, err := s.repo.GetDiscountById(id)

	if err != nil {
		return errors.New("Discount not found")
	}
	if req.Rate > 0 {
		existing.Rate = req.Rate
	}
	if req.Desc != "" {
		existing.Desc = req.Desc
	}
	if req.IsFlashSale != nil {
		existing.IsFlashSale = *req.IsFlashSale
	}

	return s.repo.Update(id, *existing)
}

func (s *DiscountService) Delete(id int) error {
	return s.repo.Delete(id)
}
