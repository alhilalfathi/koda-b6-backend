package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(rp *repository.CartRepository) *CartService {
	return &CartService{
		repo: rp,
	}
}

func (s *CartService) CreateCart(req *models.CreateCartRequest) error {
	if req.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	if req.UserId <= 0 {
		return errors.New("invalid user id: user must be logged in")
	}

	cart := models.Cart{
		Quantity:  req.Quantity,
		Size:      req.Size,
		Variant:   req.Variant,
		UserId:    req.UserId,
		ProductId: req.ProductId,
	}
	return s.repo.CreateCart(cart)
}

func (s *CartService) GetAllCarts() ([]models.Cart, error) {
	return s.repo.GetAllCarts()
}

func (s *CartService) GetCartById(id int) (*models.Cart, error) {
	return s.repo.GetCartById(id)
}

func (s *CartService) Update(id int, req models.UpdateCartRequest) error {
	existing, err := s.repo.GetCartById(id)

	if err != nil {
		return errors.New("cart not found")
	}
	if req.Quantity > 0 {
		existing.Quantity = req.Quantity
	}
	if req.Size != "" {
		existing.Size = req.Size
	}
	if req.Variant != "" {
		existing.Variant = req.Variant
	}
	if req.UserId > 0 {
		existing.UserId = req.UserId
	}
	if req.ProductId > 0 {
		existing.ProductId = req.ProductId
	}

	return s.repo.Update(id, *existing)
}

func (s *CartService) Delete(id int) error {
	return s.repo.Delete(id)
}
