package service

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type LandingService struct {
	repo *repository.ProductRepository
}

func NewLandingService(rp *repository.ProductRepository) *LandingService {
	return &LandingService{
		repo: rp,
	}
}

func (s *LandingService) RecommendedProducts() ([]models.Product, error) {
	return s.repo.RecomendedProducts()
}
