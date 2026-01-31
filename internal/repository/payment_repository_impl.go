package repository

import (
	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (r paymentRepositoryImpl) CreatePayment(payment *entity.Payment) (*model.Payment, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO payments (
			user_id,
			currency,
			address_out,
			address_in,
			status
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			"id"
			"currency"
	`

	var response model.Payment

	err := r.db.QueryRow(
		query,
		payment.UserID,
		payment.Currency,
		payment.AddressOut,
		payment.AddressIn,
		payment.Status,
	).Scan(&response.ID, &response.Currency)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *paymentRepositoryImpl) GetPaymentByID(payment *entity.Payment) (*model.Payment, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		id,
		user_id,
		amount,
		currency,
		tx_id,
		status,
		created_at,
		confirmed_at
	FROM payments
	WHERE id = $1
	`

	var response model.Payment

	err := r.db.QueryRow(
		query,
		payment.ID,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.Amount,
		&response.Currency,
		&response.TxID,
		&response.Status,
		&response.CreatedAt,
		&response.ConfirmedAt,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *paymentRepositoryImpl) GetPaymentsByUserID(payment *entity.Payment) (*[]model.Payment, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		id,
		user_id,
		amount,
		currency,
		tx_id,
		status,
		created_at,
		confirmed_at
	FROM payments
	WHERE user_id = $1
	`

	var response []model.Payment

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var payment model.Payment

		if err := rows.Scan(
			&payment.ID,
			&payment.UserID,
			&payment.Amount,
			&payment.Currency,
			&payment.TxID,
			&payment.Status,
			&payment.CreatedAt,
			&payment.ConfirmedAt,
		); err != nil {
			return nil, err
		}

		response = append(response, payment)
	}

	return &response, nil
}

func (r *paymentRepositoryImpl) DeletePaymentByID(payment *entity.Payment) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	DELETE
	FROM payments
	WHERE id = $1;
	`

	err := r.db.QueryRow(
		query,
		payment.ID,
	).Scan()

	if err != nil {
		return false, err
	}

	return true, nil
}
