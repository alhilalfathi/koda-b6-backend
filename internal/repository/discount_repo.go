package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DiscountRepository struct {
	db *pgxpool.Pool
}

func NewDiscountRepository(d *pgxpool.Pool) *DiscountRepository {
	return &DiscountRepository{
		db: d,
	}
}

func (r *DiscountRepository) CreateDiscount(dis models.Discount) error {
	query := `INSERT INTO "DISCOUNT" (discount_rate, description, is_flashsale) VALUES ($1,$2,$3)`

	_, err := r.db.Exec(context.Background(), query, dis.Rate, dis.Desc, dis.IsFlashSale)

	return err
}

func (r *DiscountRepository) GetAllDiscounts() ([]models.Discount, error) {
	query := `SELECT "id", "discount_rate", "description", "is_flashsale" FROM "DISCOUNT"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	disc, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Discount])
	if err != nil {
		return nil, err
	}

	return disc, nil
}

func (r *DiscountRepository) GetDiscountById(id int) (*models.Discount, error) {
	query := `SELECT "id", "discount_rate", "description", "is_flashsale" FROM "DISCOUNT" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	disc, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Discount])
	if err != nil {
		return nil, err
	}
	return &disc, nil
}

func (r *DiscountRepository) Update(id int, disc models.Discount) error {

	query := `UPDATE "DISCOUNT" SET "discount_rate"=$1, "description"=$2, "is_flashsale"=$3 WHERE "id"=$4`

	_, err := r.db.Exec(context.Background(), query, disc.Rate, disc.Desc, disc.IsFlashSale, id)
	return err
}

func (r *DiscountRepository) Delete(id int) error {
	query := `DELETE FROM "DISCOUNT" WHERE "id" = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}
