package service

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type TransactionProductService struct {
	repo *repository.TransactionProductRepository
}

func NewTransactionProductService(rp *repository.TransactionProductRepository) *TransactionProductService {
	return &TransactionProductService{
		repo: rp,
	}
}

func (s *TransactionProductService) Create(req *models.CreateTransactionProductRequest) error {

	product := models.TransactionProduct{
		TrId:      req.TrId,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
		Size:      req.Size,
		Variant:   req.Variant,
	}
	return s.repo.Create(product)
}

func (s *TransactionProductService) GetAll() ([]models.TransactionProduct, error) {
	return s.repo.GetAll()
}

func (s *TransactionProductService) GetById(id int) (*models.TransactionProduct, error) {
	return s.repo.GetById(id)
}
