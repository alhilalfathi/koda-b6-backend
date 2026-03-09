package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type ReviewRepository struct {
	db *pgx.Conn
}

func NewReviewRepository(d *pgx.Conn) *ReviewRepository {
	return &ReviewRepository{
		db: d,
	}
}

func (r *ReviewRepository) CreateReview(review models.Review) error {
	query := `INSERT INTO "PRODUCT" ("user_id", "product_id", "messages", "rating") VALUES ($1,$2,$3,$4)`

	_, err := r.db.Exec(context.Background(), query, review.UserId, review.ProductId, review.Messages, review.Rating)

	return err
}

func (r *ReviewRepository) GetAllReviews() ([]models.Review, error) {
	query := `SELECT "id", "user_id", "product_id", "messages", "rating" FROM "REVIEWS"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	reviews, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Review])
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepository) GetReviewById(id int) (*models.Review, error) {
	query := `SELECT "id", "user_id", "product_id", "messages", "rating" FROM "REVIEWS" WHERE "id" = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	review, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Review])
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) Update(id int, review models.Review) error {

	query := `UPDATE "REVIEWS" SET "messages"=$1, "rating"=$2 WHERE "id"=$3`

	_, err := r.db.Exec(context.Background(), query, review.Messages, review.Rating, id)
	return err
}

func (r *ReviewRepository) Delete(id int) error {
	query := `DELETE FROM "REVIEWS" WHERE "id" = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}
