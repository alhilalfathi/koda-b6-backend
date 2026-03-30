package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductCategoryRepository struct {
	db *pgxpool.Pool
}

func NewProductCategoryRepository(d *pgxpool.Pool) *ProductCategoryRepository {
	return &ProductCategoryRepository{
		db: d,
	}
}

func (r *ProductCategoryRepository) Create(p models.ProductCategory) error {
	query := `INSERT INTO "PRODUCT_CATEGORY" (product_id, category_id) VALUES ($1,$2)`

	_, err := r.db.Exec(context.Background(), query, p.ProductId, p.CategoryId)

	return err
}

func (r *ProductCategoryRepository) GetAll() ([]models.ProductCategory, error) {
	query := `SELECT "id", "product_id", "category_id" FROM "PRODUCT_CATEGORY"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	p, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductCategory])
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProductCategoryRepository) GetById(id int) (*models.ProductCategory, error) {
	query := `SELECT "id", "product_id", "category_id" FROM "PRODUCT_CATEGORY" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ProductCategory])
	if err != nil {
		return nil, err
	}
	return &p, nil
}
