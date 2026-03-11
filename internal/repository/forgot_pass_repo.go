package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type ForgotPassRepository struct {
	db *pgx.Conn
}

func NewForgotPassRepository(d *pgx.Conn) *ForgotPassRepository {
	return &ForgotPassRepository{
		db: d,
	}
}

func (r *ForgotPassRepository) CreateForgotPass(p models.ForgotPass) error {
	query := `INSERT INTO "FORGOT_PASS" (email, code) VALUES ($1,$2)`

	_, err := r.db.Exec(context.Background(), query, p.Email, p.Code)

	return err
}

func (r *ForgotPassRepository) GetByEmailCode(email string, code string) (*models.ForgotPass, error) {
	query := `SELECT id, email, code FROM "FORGOT_PASS" WHERE email=$1 AND code=$2`

	rows, err := r.db.Query(context.Background(), query, email, code)
	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPass])
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *ForgotPassRepository) Delete(code string) error {
	query := `DELETE FROM "FORGOT_PASS" WHERE "code" = $1`
	_, err := r.db.Exec(context.Background(), query, code)
	return err
}
