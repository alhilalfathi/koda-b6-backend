package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SizeRepository struct {
	db *pgxpool.Pool
}

func NewSizeRepository(d *pgxpool.Pool) *SizeRepository {
	return &SizeRepository{
		db: d,
	}
}

func (r *SizeRepository) CreateSize(size models.Size) error {
	query := `INSERT INTO "SIZE" (size, add_price) VALUES ($1,$2)`

	_, err := r.db.Exec(context.Background(), query, size.Size, size.AddedPrice)

	return err
}

func (r *SizeRepository) GetAllSizes() ([]models.Size, error) {
	query := `SELECT "id", "size", "add_price" FROM "SIZE"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	sizes, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Size])
	if err != nil {
		return nil, err
	}

	return sizes, nil
}

func (r *SizeRepository) GetSizeById(id int) (*models.Size, error) {
	query := `SELECT "id", "size", "add_price" FROM "SIZE" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	size, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Size])
	if err != nil {
		return nil, err
	}
	return &size, nil
}

func (r *SizeRepository) Update(id int, size models.Size) error {

	query := `UPDATE "SIZE" SET "size"=$1, "add_price"=$2 WHERE "id"=$3`

	_, err := r.db.Exec(context.Background(), query, size.Size, size.AddedPrice, id)
	return err
}

func (r *SizeRepository) Delete(id int) error {
	query := `DELETE FROM "SIZE" WHERE "id" = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}
