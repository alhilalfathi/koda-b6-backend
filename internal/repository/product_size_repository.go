package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductSizeRepository struct {
	db *pgxpool.Pool
}

func NewProductSizeRepository(d *pgxpool.Pool) *ProductSizeRepository {
	return &ProductSizeRepository{
		db: d,
	}
}

func (r *ProductSizeRepository) Create(p models.ProductSize) error {
	query := `INSERT INTO "PRODUCT_size" (product_id, size_id) VALUES ($1,$2)`

	_, err := r.db.Exec(context.Background(), query, p.ProductId, p.SizeId)

	return err
}

func (r *ProductSizeRepository) GetAll() ([]models.ProductSize, error) {
	query := `SELECT "id", "product_id", "size_id" FROM "PRODUCT_size"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	p, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductSize])
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProductSizeRepository) GetById(id int) (*models.ProductSize, error) {
	query := `SELECT "id", "product_id", "size_id" FROM "PRODUCT_size" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ProductSize])
	if err != nil {
		return nil, err
	}
	return &p, nil
}
