package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type SizeService struct {
	repo *repository.SizeRepository
}

func NewSizeService(rp *repository.SizeRepository) *SizeService {
	return &SizeService{
		repo: rp,
	}
}

func (s *SizeService) CreateSize(req *models.CreateSizeRequest) error {
	if req.Size != "" {
		return errors.New("Input Invalid")
	}
	if req.AddedPrice < 0 {
		return errors.New("Input Invalid")
	}

	size := models.Size{
		Size:       req.Size,
		AddedPrice: req.AddedPrice,
	}
	return s.repo.CreateSize(size)
}

func (s *SizeService) GetAllSizes() ([]models.Size, error) {
	return s.repo.GetAllSizes()
}

func (s *SizeService) GetSizeById(id int) (*models.Size, error) {
	return s.repo.GetSizeById(id)
}

func (s *SizeService) Update(id int, req models.UpdateSizeRequest) error {
	existing, err := s.repo.GetSizeById(id)

	if err != nil {
		return errors.New("Size not found")
	}
	if req.Size != "" {
		existing.Size = req.Size
	}

	return s.repo.Update(id, *existing)
}

func (s *SizeService) Delete(id int) error {
	return s.repo.Delete(id)
}
