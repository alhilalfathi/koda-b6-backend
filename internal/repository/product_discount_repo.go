package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type ProductDiscountRepository struct {
	db *pgx.Conn
}

func NewProductDiscountRepository(d *pgx.Conn) *ProductDiscountRepository {
	return &ProductDiscountRepository{
		db: d,
	}
}

func (r *ProductDiscountRepository) Create(p models.ProductDiscount) error {
	query := `INSERT INTO "PRODUCT_DISCOUNT" (product_id, discount_id) VALUES ($1,$2)`

	_, err := r.db.Exec(context.Background(), query, p.ProductId, p.DiscountId)

	return err
}

func (r *ProductDiscountRepository) GetAll() ([]models.ProductDiscount, error) {
	query := `SELECT "id", "product_id", "discount_id" FROM "PRODUCT_DISCOUNT"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	p, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductDiscount])
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProductDiscountRepository) GetById(id int) (*models.ProductDiscount, error) {
	query := `SELECT "id", "product_id", "discount_id" FROM "PRODUCT_DISCOUNT" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ProductDiscount])
	if err != nil {
		return nil, err
	}
	return &p, nil
}
