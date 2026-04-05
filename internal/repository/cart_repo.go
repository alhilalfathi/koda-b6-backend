package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CartRepository struct {
	db *pgxpool.Pool
}

func NewCartRepository(d *pgxpool.Pool) *CartRepository {
	return &CartRepository{
		db: d,
	}
}

func (r *CartRepository) CreateCart(cart models.Cart) error {
	query := `
        INSERT INTO "CART" (quantity, size, variant, user_id, product_id) 
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (user_id, product_id, size, variant) 
        DO UPDATE SET quantity = "CART".quantity + EXCLUDED.quantity
		`

	_, err := r.db.Exec(context.Background(), query, cart.Quantity, cart.Size, cart.Variant, cart.UserId, cart.ProductId)

	return err
}

func (r *CartRepository) GetAllCarts() ([]models.Cart, error) {
	query := `SELECT "id", "quantity", "size", "variant", "user_id", "product_id" FROM "CART"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	cart, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Cart])
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *CartRepository) GetCartById(id int) (*models.Cart, error) {
	query := `
	SELECT "id", "quantity", "size", "variant", "user_id", "product_id" 
	FROM "CART" 
	WHERE "id" = $1
	`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	cart, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Cart])
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) GetDetailCartByUserId(userId int) ([]models.CartResponse, error) {
	query := `
    SELECT 
        c."id", 
        c."quantity", 
        c."size", 
        c."variant", 
        c."user_id", 
        c."product_id",
        p."product_name",
        p."price",
        pi."path"
    FROM "CART" c
    JOIN "PRODUCT" p ON c."product_id" = p."id"
	LEFT JOIN "PRODUCT_IMAGES" pi ON c."product_id" = pi."product_id"
    WHERE c."user_id" = $1
    `

	rows, err := r.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cartItems, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.CartResponse])
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (r *CartRepository) Update(id int, cart models.Cart) error {

	query := `UPDATE "CART" SET "quantity"=$1, "size"=$2, "variant"=$3, "user_id"=$4, "product_id"=$5 WHERE "id"=$6`

	_, err := r.db.Exec(context.Background(), query, cart.Quantity, cart.Size, cart.Variant, cart.UserId, cart.ProductId, id)
	return err
}

func (r *CartRepository) Delete(id int) error {
	query := `DELETE FROM "CART" WHERE "id" = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}
