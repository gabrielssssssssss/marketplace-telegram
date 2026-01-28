package repository

import (
	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (r orderRepositoryImpl) CreateOrder(orders *entity.Orders) (*model.Orders, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO orders (
			user_id,
			product_name,
			amount
		)
		VALUES ($1, $2, $3)
		RETURNING
			"id"
	`

	var response model.Orders

	err := r.db.QueryRow(
		query,
		orders.UserID,
		orders.ProductName,
		orders.Amount,
	).Scan(&response.ID)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *orderRepositoryImpl) GetOrderByID(orders *entity.Orders) (*model.Orders, error) {
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

	var response model.Orders

	err := r.db.QueryRow(
		query,
		orders.ID,
		orders.UserID,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.ProductName,
		&response.Amount,
		&response.OrderAt,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *orderRepositoryImpl) GetOrdersByUserID(orders *entity.Orders) (*[]model.Orders, error) {
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

	var response []model.Orders

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Orders

		if err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.ProductName,
			&order.Amount,
			&order.OrderAt,
		); err != nil {
			return nil, err
		}

		response = append(response, order)
	}

	return &response, nil
}

func (r *orderRepositoryImpl) DeleteOrderByID(orders *entity.Orders) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	DELETE
	FROM orders
	WHERE id = $1;
	`

	err := r.db.QueryRow(
		query,
		orders.ID,
	).Scan()

	if err != nil {
		return false, err
	}

	return true, nil
}
