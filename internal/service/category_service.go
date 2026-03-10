package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(rp *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: rp,
	}
}

func (s *CategoryService) CreateCategory(req *models.CreateCategoryRequest) error {
	if req.Category != "" {
		return errors.New("Input invalid")
	}

	cat := models.Category{
		Category: req.Category,
	}
	return s.repo.CreateCategory(cat)
}

func (s *CategoryService) GetAllCategory() ([]models.Category, error) {
	return s.repo.GetAllCategory()
}

func (s *CategoryService) GetCategoryById(id int) (*models.Category, error) {
	return s.repo.GetCategoryById(id)
}

func (s *CategoryService) Update(id int, req models.UpdateCategoryRequest) error {
	existing, err := s.repo.GetCategoryById(id)

	if err != nil {
		return errors.New("category not found")
	}
	if req.Category != "" {
		existing.Category = req.Category
	}

	return s.repo.Update(id, *existing)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
