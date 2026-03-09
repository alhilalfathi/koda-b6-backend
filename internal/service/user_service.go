package service

import (
	"errors"
	"fmt"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"

	"github.com/matthewhartstonge/argon2"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(rp *repository.UserRepository) *UserService {
	return &UserService{
		repo: rp,
	}
}

func (s *UserService) GetAll() ([]models.Users, error) {
	users, err := s.repo.GetAllUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetById(id string) (*models.Users, error) {

	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByEmail(email string) (*models.Users, error) {

	user, err := s.repo.GetById(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Register(req *models.CreateUserRequest) error {
	existingUser, _ := s.repo.GetByEmail(req.Email)

	if existingUser != nil {
		return errors.New("Email is registered")
	}

	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(req.Password))

	if err != nil {
		return err
	}

	newUser := models.Users{
		FullName: req.FullName,
		Email:    req.Email,
		Password: string(encoded),
	}
	return s.repo.Create(newUser)
}

func (s *UserService) Create(req *models.CreateUserRequest) error {

	if req.Email == "" || req.Password == "" {
		return errors.New("Email and password required")
	}

	newUser := models.Users{
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.repo.Create(newUser)
	if err != nil {
		return fmt.Errorf("Failed to create user: %w", err)
	}

	return nil
}

func (s *UserService) Update(email string, u *models.UpdateUserRequest) (*models.Users, error) {

	if u.Password == "" {
		return nil, errors.New("Password cannot blank")
	}

	user := &models.Users{
		Password: u.Password,
	}

	updatedUser, err := s.repo.Update(email, user)
	if err != nil {
		return nil, fmt.Errorf("Failed to update user: %w", err)
	}

	return updatedUser, nil
}

func (s *UserService) Delete(id string) error {

	_, err := s.repo.GetById(id)
	if err != nil {
		return errors.New("User not found")
	}

	err = s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("Failed to delete user: %w", err)
	}

	return nil
}
