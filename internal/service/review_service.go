package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ReviewService struct {
	repo *repository.ReviewRepository
}

func NewReviewService(rp *repository.ReviewRepository) *ReviewService {
	return &ReviewService{
		repo: rp,
	}
}

func (s *ReviewService) CreateReview(req *models.CreateReviewRequest) error {

	review := models.Review{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Messages:  req.Messages,
		Rating:    req.Rating,
	}
	return s.repo.CreateReview(review)
}

func (s *ReviewService) GetAllReviews() ([]models.Review, error) {
	return s.repo.GetAllReviews()
}

func (s *ReviewService) GetReviewById(id int) (*models.Review, error) {
	return s.repo.GetReviewById(id)
}

func (s *ReviewService) Update(id int, req models.UpdateReviewRequest) error {
	existing, err := s.repo.GetReviewById(id)

	if err != nil {
		return errors.New("product not found")
	}
	if req.Messages != "" {
		existing.Messages = req.Messages
	}
	if req.Rating > 0 {
		existing.Rating = req.Rating
	}

	return s.repo.Update(id, *existing)
}

func (s *ReviewService) Delete(id int) error {
	return s.repo.Delete(id)
}
