package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(d *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{
		db: d,
	}
}

func (r *CategoryRepository) CreateCategory(cat models.Category) error {
	query := `INSERT INTO "CATEGORY" (category) VALUES ($1)`

	_, err := r.db.Exec(context.Background(), query, cat.Category)

	return err
}

func (r *CategoryRepository) GetAllCategory() ([]models.Category, error) {
	query := `SELECT "id", "category" FROM "CATEGORY"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	cat, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Category])
	if err != nil {
		return nil, err
	}

	return cat, nil
}

func (r *CategoryRepository) GetCategoryById(id int) (*models.Category, error) {
	query := `SELECT "id", "category" FROM "CATEGORY" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	cat, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Category])
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *CategoryRepository) Update(id int, cat models.Category) error {

	query := `UPDATE "CATEGORY" SET "category"=$1 WHERE "id"=$2`

	_, err := r.db.Exec(context.Background(), query, cat.Category, id)
	return err
}

func (r *CategoryRepository) Delete(id int) error {
	query := `DELETE FROM "CATEGORY" WHERE "id" = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}
