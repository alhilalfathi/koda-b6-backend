package repository

import (
	"context"
	"koda-b6-backend/internal/models"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(d *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: d,
	}
}

func (r *UserRepository) GetAllUser() ([]models.Users, error) {
	query := `SELECT id, email, password FROM "USER"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetById(id string) (*models.Users, error) {
	numId, _ := strconv.Atoi(id)

	query := `SELECT id, fullname, email, password FROM "USER" WHERE id=$1`

	rows, err := r.db.Query(context.Background(), query, numId)
	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.Users, error) {
	query := `SELECT id, fullname, email, password FROM "USER" WHERE email=$1`

	rows, err := r.db.Query(context.Background(), query, email)
	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user models.Users) error {
	query := `INSERT INTO "USER" (fullname, email, password) VALUES ($1,$2,$3)`

	_, err := r.db.Exec(context.Background(), query, user.FullName, user.Email, user.Password)

	return err
}

func (r *UserRepository) Update(email string, user *models.Users) (*models.Users, error) {
	query := `UPDATE "USER"	SET password=$1	WHERE email=$2	RETURNING id, email, password`

	rows, err := r.db.Query(context.Background(), query, user.Password, email)
	if err != nil {
		return nil, err
	}

	updatedUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (r *UserRepository) Delete(email string) error {
	query := `DELETE FROM "USER" WHERE email=$1`

	_, err := r.db.Exec(context.Background(), query, email)

	return err
}
