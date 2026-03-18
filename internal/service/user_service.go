package service

import (
	"errors"
	"koda-b6-backend/internal/lib"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
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
	return s.repo.GetAllUser()
}

func (s *UserService) GetById(id string) (*models.Users, error) {

	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByEmail(email string) (*models.Users, error) {

	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Register(req *models.CreateUserRequest) error {
	existingUser, _ := s.repo.GetByEmail(req.Email)

	if existingUser != nil {
		return errors.New("Email was registered")
	}

	hashed, err := lib.HashPassword(req.Password)
	if err != nil {
		return err
	}

	newUser := models.Users{
		FullName: req.FullName,
		Email:    req.Email,
		Password: hashed,
	}
	return s.repo.Create(newUser)
}

func (s *UserService) Login(req models.LoginUserRequest) (*models.Users, string, error) {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("Invalid email or password")
	}

	ok := lib.VerifyPassword(req.Password, user.Password)
	if !ok {
		return nil, "", errors.New("Invalid Email or Password")
	}

	if ok {
		token, err := lib.GenerateToken(user.Id)
		if err != nil {
			return nil, "", err
		}

		return user, token, nil
	}
	return nil, "", err
}

func (s *UserService) Update(email string, req *models.UpdateUserRequest) (*models.Users, error) {
	user, err := s.repo.GetByEmail("email")
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password
	}

	return s.repo.Update(email, user)
}

func (s *UserService) Delete(email string) error {
	_, err := s.repo.GetByEmail(email)
	if err != nil {
		return errors.New("User not found")
	}

	return s.repo.Delete(email)
}
