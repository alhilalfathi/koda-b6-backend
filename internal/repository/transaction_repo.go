package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type TransactionRepository struct {
	db *pgx.Conn
}

func NewTransactionRepository(d *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{
		db: d,
	}
}

func (r *TransactionRepository) CreateTransaction(tr models.Transaction) error {
	query := `
		INSERT INTO "TRANSACTION" (
			trx_id, 
			user_id,  
			fullname, 
			email, 
			address, 
			delivery, 
			delivery_fee, 
			tax, 
			total, 
			status_order
			) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	_, err := r.db.Exec(context.Background(), query,
		tr.TrxId,
		tr.UserId,
		tr.Fullname,
		tr.Email,
		tr.Address,
		tr.Delivery,
		tr.DeliveryFee,
		tr.Tax,
		tr.Total,
		tr.OrderStatus)

	return err
}

func (r *TransactionRepository) GetAllTransaction() ([]models.Transaction, error) {
	query := `
	SELECT "id", 
			"trx_id", 
			"user_id", 
			"order_date", 
			"fullname", 
			"email", 
			"address", 
			"delivery", 
			"delivery_fee", 
			"tax", 
			"total", 
			"status_order"
	FROM "PRODUCT" 
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	tr, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Transaction])
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func (r *TransactionRepository) GetTransactionById(id int) (*models.Transaction, error) {
	query := `
	SELECT "id", 
			"trx_id", 
			"user_id", 
			"order_date", 
			"fullname", 
			"email", 
			"address", 
			"delivery", 
			"delivery_fee", 
			"tax", 
			"total", 
			"status_order"
	FROM "PRODUCT" 
	WHERE "id" = $1
	`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	tr, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Transaction])
	if err != nil {
		return nil, err
	}
	return &tr, nil
}
