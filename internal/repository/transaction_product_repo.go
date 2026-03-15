package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type TransactionProductRepository struct {
	db *pgx.Conn
}

func NewTransactionProductRepository(d *pgx.Conn) *TransactionProductRepository {
	return &TransactionProductRepository{
		db: d,
	}
}

func (r *TransactionProductRepository) Create(p models.TransactionProduct) error {
	query := `INSERT INTO "TRANSACTION_PRODUCT" (transaction_id, product_id, quantity, size, variant) VALUES ($1,$2,$3,$4,$5)`

	_, err := r.db.Exec(context.Background(), query, p.TrId, p.ProductId, p.Quantity, p.Size, p.Variant)

	return err
}

func (r *TransactionProductRepository) GetAll() ([]models.TransactionProduct, error) {
	query := `SELECT "id", "transaction_id", "product_id", "quantity", "size", "variant_id" FROM "TRANSACTION_PRODUCT"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	p, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.TransactionProduct])
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *TransactionProductRepository) GetById(id int) (*models.TransactionProduct, error) {
	query := `SELECT "id", "transaction_id", "product_id", "quantity", "size", "variant_id" FROM "TRANSACTION_PRODUCT" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.TransactionProduct])
	if err != nil {
		return nil, err
	}
	return &p, nil
}
