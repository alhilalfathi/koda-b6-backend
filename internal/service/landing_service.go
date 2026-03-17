package service

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type LandingService struct {
	repo     *repository.ProductRepository
	rpReview *repository.ReviewRepository
}

func NewLandingService(rp *repository.ProductRepository, rv *repository.ReviewRepository) *LandingService {
	return &LandingService{
		repo:     rp,
		rpReview: rv,
	}
}

func (s *LandingService) RecommendedProducts() ([]models.RecommendedProduct, error) {
	return s.repo.RecomendedProducts()
}

func (s *LandingService) GetAllReviews() ([]models.Review, error) {
	return s.rpReview.GetAllReviews()
}
