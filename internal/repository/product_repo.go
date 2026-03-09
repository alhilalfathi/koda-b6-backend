package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(d *pgx.Conn) *ProductRepository {
	return &ProductRepository{
		db: d,
	}
}

func (r *ProductRepository) CreateProduct(product models.Product) error {
	query := `INSERT INTO "PRODUCT" (product_name, product_desc, price, stock) VALUES ($1,$2,$3,$4)`

	_, err := r.db.Exec(context.Background(), query, product.ProductName, product.Desc, product.Price, product.Stock)

	return err
}

func (r *ProductRepository) GetAllProduct() ([]models.Product, error) {
	query := `SELECT "id", "product_name", "product_desc", "price", "stock" FROM "PRODUCT"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product

		err := rows.Scan(
			&p.Id,
			&p.ProductName,
			&p.Desc,
			&p.Price,
			&p.Stock,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) GetProductById(id int) (*models.Product, error) {
	query := `SELECT "id", "product_name", "product_desc", "price", "stock" FROM "PRODUCT" WHERE "id" = $1`

	var p models.Product

	err := r.db.QueryRow(context.Background(), query, id).
		Scan(&p.Id, &p.ProductName, &p.Desc, &p.Price, &p.Stock)

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Update(id int, p models.Product) error {

	query := `UPDATE "PRODUCT" SET "product_name"=$1, "product_desc"=$2, "price"=$3, "stock"=$4 WHERE "id"=$5`

	_, err := r.db.Exec(context.Background(), query, p.ProductName, p.Desc, p.Price, p.Stock, id)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	query := `DELETE FROM "PRODUCT" WHERE "id" = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}
