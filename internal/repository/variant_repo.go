package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type VariantRepository struct {
	db *pgx.Conn
}

func NewVariantRepository(d *pgx.Conn) *VariantRepository {
	return &VariantRepository{
		db: d,
	}
}

func (r *VariantRepository) CreateVariant(variant models.Variant) error {
	query := `INSERT INTO "VARIANT" (variant, add_price) VALUES ($1,$2)`

	_, err := r.db.Exec(context.Background(), query, variant.Variant, variant.AddedPrice)

	return err
}

func (r *VariantRepository) GetAllVariants() ([]models.Variant, error) {
	query := `SELECT "id", "variant", "add_price" FROM "VARIANT"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	variant, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Variant])
	if err != nil {
		return nil, err
	}

	return variant, nil
}

func (r *VariantRepository) GetVariantById(id int) (*models.Variant, error) {
	query := `SELECT "id", "variant", "add_price" FROM "VARIANT" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	variant, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Variant])
	if err != nil {
		return nil, err
	}
	return &variant, nil
}

func (r *VariantRepository) Update(id int, variant models.Variant) error {

	query := `UPDATE "VARIANT" SET "variant"=$1, "add_price"=$2 WHERE "id"=$3`

	_, err := r.db.Exec(context.Background(), query, variant.Variant, variant.AddedPrice, id)
	return err
}

func (r *VariantRepository) Delete(id int) error {
	query := `DELETE FROM "VARIANT" WHERE "id" = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}
