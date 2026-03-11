package service

import (
	"errors"
	"fmt"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
	"math/rand"
)

type ForgotPassService struct {
	forgotPassRepo *repository.ForgotPassRepository
	userRepo       *repository.UserRepository
}

func NewForgotPassService(fr *repository.ForgotPassRepository, ur *repository.UserRepository) *ForgotPassService {
	return &ForgotPassService{
		forgotPassRepo: fr,
		userRepo:       ur,
	}
}

func (s *ForgotPassService) RequestForgotPass(req *models.CreateForgotPassRequest) error {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return errors.New("Email not found")
	}

	code := rand.Intn(99999)
	fmt.Println(code)

	result := models.ForgotPass{
		Email: user.Email,
		Code:  code,
	}

	return s.forgotPassRepo.CreateForgotPass(result)
}

func (s *ForgotPassService) ResetPass(req *models.CreateForgotPassRequest) error {
	_, err := s.forgotPassRepo.GetByEmailCode(req.Email, req.Code)
	if err != nil {
		return errors.New("Email and Code invalid")
	}

	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return errors.New("Email not found")
	}

	s.userRepo.Update(req.Email, user)

	s.forgotPassRepo.Delete(req.Code)

	return nil
}
