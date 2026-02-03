package repository

import (
	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (r cartRepositoryImpl) CreateCart(cart *entity.Cart) (*model.Cart, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO cart (
			user_id,
			product_id
		)
		VALUES ($1, $2)
		RETURNING
			id,
			user_id,
			product_id,
			created_at,
			updated_at
	`

	var response model.Cart

	err := r.db.QueryRow(
		query,
		cart.UserID,
		cart.ProductID,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.ProductID,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	return &response, err
}

func (r cartRepositoryImpl) GetCartByID(cart *entity.Cart) (*model.Cart, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		id,
		user_id,
		product_id,
		created_at,
		updated_at
	FROM cart
	WHERE id = $1
	`

	var response model.Cart

	err := r.db.QueryRow(
		query,
		cart.ID,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.ProductID,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	return &response, err
}

func (r cartRepositoryImpl) GetCartsByUserID(cart *entity.Cart) (*[]model.Cart, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		id,
		user_id,
		product_id,
		created_at,
		updated_at
	FROM cart
	WHERE user_id = $1
	`

	var response []model.Cart

	rows, err := r.db.Query(query, cart.UserID)
	if err != nil {
		return &response, err
	}
	defer rows.Close()

	for rows.Next() {
		var row model.Cart

		if err := rows.Scan(
			&row.ID,
			&row.UserID,
			&row.ProductID,
			&row.CreatedAt,
			&row.UpdatedAt,
		); err != nil {
			return nil, err
		}
		response = append(response, row)
	}

	return &response, err
}

func (r *cartRepositoryImpl) UpdateCartByID(cart *entity.Cart) (*model.Cart, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
    UPDATE cart
    SET 
		user_id      = COALESCE($1, user_id),
		product_id   = COALESCE($2, product_id),
		updated_at   = COALESCE($3, updated_at)
    WHERE id = $4
    RETURNING
		id,
		user_id,
		product_id,
		created_at,
		updated_at
    `

	var response model.Cart

	err := r.db.QueryRow(
		query,
		cart.UserID,
		cart.ProductID,
		cart.UpdatedAt,
		cart.ID,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.ProductID,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	return &response, err
}

func (r cartRepositoryImpl) DeleteCartByID(cart *entity.Cart) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	_, err := r.db.Query(`DELETE FROM cart WHERE id = $1;`, cart.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
