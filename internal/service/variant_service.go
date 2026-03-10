package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type VariantService struct {
	repo *repository.VariantRepository
}

func NewVariantService(rp *repository.VariantRepository) *VariantService {
	return &VariantService{
		repo: rp,
	}
}

func (s *VariantService) CreateVariant(req *models.CreateVariantRequest) error {
	if req.Variant != "" {
		return errors.New("Input Invalid")
	}
	if req.AddedPrice < 0 {
		return errors.New("Input Invalid")
	}

	variant := models.Variant{
		Variant:    req.Variant,
		AddedPrice: req.AddedPrice,
	}
	return s.repo.CreateVariant(variant)
}

func (s *VariantService) GetAllVariants() ([]models.Variant, error) {
	return s.repo.GetAllVariants()
}

func (s *VariantService) GetVariantById(id int) (*models.Variant, error) {
	return s.repo.GetVariantById(id)
}

func (s *VariantService) Update(id int, req models.UpdateVariantRequest) error {
	existing, err := s.repo.GetVariantById(id)

	if err != nil {
		return errors.New("Variant not found")
	}
	if req.Variant != "" {
		existing.Variant = req.Variant
	}

	return s.repo.Update(id, *existing)
}

func (s *VariantService) Delete(id int) error {
	return s.repo.Delete(id)
}
