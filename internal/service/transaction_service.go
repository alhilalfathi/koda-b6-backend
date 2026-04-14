package service

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(rp *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: rp,
	}
}

func (s *TransactionService) CreateTransaction(req *models.CreateTransactionRequest) error {

	product := models.Transaction{
		TrxId:       req.TrxId,
		UserId:      req.UserId,
		Fullname:    req.Fullname,
		Email:       req.Email,
		Address:     req.Address,
		Delivery:    req.Delivery,
		DeliveryFee: req.DeliveryFee,
		Tax:         req.Tax,
		Total:       req.Total,
		OrderStatus: req.OrderStatus,
	}
	return s.repo.CreateTransaction(product)
}

func (s *TransactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repo.GetAllTransaction()
}

func (s *TransactionService) GetDetail(id int) (*models.TransactionDetailResponse, error) {
	return s.repo.GetDetail(id)
}
