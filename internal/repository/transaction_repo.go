package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(d *pgxpool.Pool) *TransactionRepository {
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

func (r *TransactionRepository) GetDetail(id int) (*models.TransactionDetailResponse, error) {
	query := `
	SELECT 
		t.id,
		t.trx_id,
		t.order_date,
		t.total,
		t.status_order,

		tp.product_id,
		tp.qty,

		p.product_name,
		p.price,
		COALESCE(pi.path, '') as path

	FROM "TRANSACTION" t
	JOIN "TRANSACTION_PRODUCT" tp ON tp.transaction_id = t.id
	JOIN "PRODUCT" p ON p.id = tp.product_id
	LEFT JOIN "PRODUCT_IMAGE" pi ON pi.product_id = p.id

	WHERE t.id = $1
	`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result models.TransactionDetailResponse
	items := []models.TransactionItem{}

	for rows.Next() {
		var item models.TransactionItem

		err := rows.Scan(
			&result.Id,
			&result.TrxId,
			&result.OrderDate,
			&result.Total,
			&result.StatusOrder,
			&item.ProductId,
			&item.Qty,
			&item.ProductName,
			&item.Price,
			&item.Image,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	result.Items = items

	return &result, nil
}
