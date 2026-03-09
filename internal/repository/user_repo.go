package repository

import (
	"context"
	"koda-b6-backend/internal/models"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(d *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: d,
	}
}

func (r *UserRepository) GetAllUser() ([]models.Users, error) {

	rows, err := r.db.Query(context.Background(), `SELECT id,email,password FROM "USER"`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.Users

	for rows.Next() {
		var user models.Users

		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetById(id string) (*models.Users, error) {

	var user models.Users
	numId, _ := strconv.Atoi(id)

	query := `SELECT id, email, password FROM "USER" WHERE id=$1`

	err := r.db.QueryRow(context.Background(), query, numId).
		Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Register(user models.Users) {

}

func (r *UserRepository) Create(user models.Users) error {

	query := `INSERT INTO "USER" (email, password) VALUES ($1,$2)`

	_, err := r.db.Exec(context.Background(), query, user.Email, user.Password)

	return err
}

func (r *UserRepository) Update(email string, user *models.Users) (*models.Users, error) {

	query := `UPDATE "USER"	SET password=$1	WHERE email=$2	RETURNING id, email, password`

	var updatedUser models.Users

	err := r.db.QueryRow(context.Background(), query, user.Password, email).
		Scan(&updatedUser.Id, &updatedUser.Email, &updatedUser.Password)

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (r *UserRepository) Delete(email string) error {

	query := `DELETE FROM "USER" WHERE email=$1`

	_, err := r.db.Exec(context.Background(), query, email)

	if err != nil {
		return err
	}

	return nil
}
