package repository

import (
	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (r orderRepositoryImpl) CreateOrder(order *entity.Order) (*model.Order, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO orders (
			user_id,
			amount
		)
		VALUES ($1, $2)
		RETURNING
			"id"
	`

	var response model.Order

	err := r.db.QueryRow(
		query,
		order.UserID,
		order.Amount,
	).Scan(&response.ID)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *orderRepositoryImpl) GetOrderByID(order *entity.Order) (*model.Order, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		id,
		user_id,
		product_name,
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

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *orderRepositoryImpl) GetOrdersByUserID(orders *entity.Order) (*[]model.Order, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		id,
		user_id,
		product_name,
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

	query := `
	DELETE
	FROM orders
	WHERE id = $1;
	`

	err := r.db.QueryRow(
		query,
		order.ID,
	).Scan()

	if err != nil {
		return false, err
	}

	return true, nil
}
