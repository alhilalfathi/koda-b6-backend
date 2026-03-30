package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductVariantRepository struct {
	db *pgxpool.Pool
}

func NewProductVariantRepository(d *pgxpool.Pool) *ProductVariantRepository {
	return &ProductVariantRepository{
		db: d,
	}
}

func (r *ProductVariantRepository) Create(p models.ProductVariant) error {
	query := `INSERT INTO "PRODUCT_VARIANT" (product_id, variant_id) VALUES ($1,$2)`

	_, err := r.db.Exec(context.Background(), query, p.ProductId, p.VariantId)

	return err
}

func (r *ProductVariantRepository) GetAll() ([]models.ProductVariant, error) {
	query := `SELECT "id", "product_id", "variant_id" FROM "PRODUCT_VARIANT"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	p, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductVariant])
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProductVariantRepository) GetById(id int) (*models.ProductVariant, error) {
	query := `SELECT "id", "product_id", "variant_id" FROM "PRODUCT_VARIANT" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ProductVariant])
	if err != nil {
		return nil, err
	}
	return &p, nil
}
