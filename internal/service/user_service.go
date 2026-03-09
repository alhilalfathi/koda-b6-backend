package service

import (
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
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
	return s.repo.GetAllUser()
}

func (s *UserService) GetById(id string) (*models.Users, error) {

	user, err := s.repo.GetById(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
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

func (s *UserService) Login(req models.LoginUserRequest) (string, error) {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	ok, err := argon2.VerifyEncoded([]byte(req.Password), []byte(user.Password))
	if err != nil || !ok {
		return "", errors.New("Invalid Email or Password")
	}

	claims := jwt.MapClaims{
		"user_id": user.Id,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("APP_SECRET")
	return token.SignedString([]byte(secret))
}

func (s *UserService) Update(email string, u *models.UpdateUserRequest) (*models.Users, error) {
	user, err := s.repo.GetByEmail("email")
	if err != nil {
		return nil, errors.New("user not found")
	}
	if u.Password == "" {
		return nil, errors.New("Password cannot blank")
	}
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(u.Password))
	if err != nil {
		return nil, err
	}

	user.Password = string(encoded)

	return s.repo.Update(email, user)
}

func (s *UserService) Delete(email string) error {
	_, err := s.repo.GetByEmail(email)
	if err != nil {
		return errors.New("User not found")
	}

	return s.repo.Delete(email)
}
