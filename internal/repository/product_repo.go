package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(d *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db: d,
	}
}

func (r *ProductRepository) CreateProduct(product models.Product) error {
	query := `INSERT INTO "PRODUCT" (product_name, product_desc, price, stock) VALUES ($1,$2,$3,$4)`

	_, err := r.db.Exec(context.Background(), query, product.ProductName, product.Desc, product.Price, product.Stock)

	return err
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := `
        SELECT 
            p."id", 
            p."product_name", 
            p."product_desc", 
            p."price", 
            p."stock", 
            MIN(img."path") AS "path", 
			string_agg( c."category" , ',' ) AS "category"
        FROM "PRODUCT" p
        LEFT JOIN "PRODUCT_IMAGES" img ON img."product_id" = p."id"
        JOIN "PRODUCT_CATEGORY" pcat ON pcat."product_id" = p."id"
        JOIN "CATEGORY" c ON c."id" = pcat."category_id"
		GROUP BY p."id", p."product_name", p."product_desc", p."price", p."stock"
    `

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductById(id int) (*models.Product, error) {
	query := `SELECT "id", "product_name", "product_desc", "price", "stock" FROM "PRODUCT" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}
	return &product, nil
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

func (r *ProductRepository) RecomendedProducts() ([]models.RecommendedProduct, error) {
	query := `
		SELECT "PRODUCT"."id", "product_name", "product_desc", "price", "stock", "path", COUNT("REVIEWS"."product_id") AS "total_review"
		FROM "PRODUCT"
		JOIN "REVIEWS" ON "REVIEWS"."product_id" = "PRODUCT"."id"
		JOIN "PRODUCT_IMAGES" ON "PRODUCT_IMAGES"."product_id" = "PRODUCT"."id"
		GROUP BY "PRODUCT"."id", "PRODUCT_IMAGES"."path"
		ORDER BY "total_review" DESC
		LIMIT 4
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.RecommendedProduct])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) UpdateImage(productId int, path string) error {
	query := `
        UPDATE "PRODUCT"
        SET path = $1
        WHERE id = $2
    `

	_, err := r.db.Exec(context.Background(), query, path, productId)
	return err
}
