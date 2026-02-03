package repository

import (
	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (r orderRepositoryImpl) InsertOrder(order *entity.Order) (*model.Order, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO orders (
			user_id,
			product,
			amount
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			user_id,
			product,
			amount,
			order_at
	`

	var response model.Order

	err := r.db.QueryRow(
		query,
		order.UserID,
		order.Product,
		order.Amount,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.Product,
		&response.Amount,
		&response.OrderAt,
	)

	return &response, err
}

func (r *orderRepositoryImpl) SelectOrderByID(order *entity.Order) (*model.Order, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		id,
		user_id,
		product,
		amount,
		order_at
	FROM orders
	WHERE id = $1
	AND user_id = $2
	`

	var response model.Order

	err := r.db.QueryRow(
		query,
		order.ID,
		order.UserID,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.Product,
		&response.Amount,
		&response.OrderAt,
	)

	return &response, err
}

func (r *orderRepositoryImpl) SelectOrdersByUserID(orders *entity.Order) (*[]model.Order, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		id,
		user_id,
		product,
		amount,
		order_at
	FROM orders
	WHERE user_id = $1
	`

	var response []model.Order

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Order

		if err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.Product,
			&order.Amount,
			&order.OrderAt,
		); err != nil {
			return nil, err
		}

		response = append(response, order)
	}

	return &response, nil
}

func (r *orderRepositoryImpl) DeleteOrderByID(order *entity.Order) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	err := r.db.QueryRow(
		`DELETE FROM orders WHERE id = $1;`,
		order.ID,
	).Scan()

	if err != nil {
		return false, err
	}

	return true, nil
}
