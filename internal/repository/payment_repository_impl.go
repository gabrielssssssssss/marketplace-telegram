package repository

import (
	"fmt"

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
			currency
		)
		VALUES ($1, $2)
		RETURNING
			"id",
			"currency"
	`

	var response model.Payment

	err := r.db.QueryRow(
		query,
		payment.UserID,
		payment.Currency,
	).Scan(&response.ID, &response.Currency)

	if err != nil {
		fmt.Println(err)
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
		currency,
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
		&response.Currency,
		&response.Status,
		&response.CreatedAt,
		&response.ConfirmedAt,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *paymentRepositoryImpl) UpdatePaymentByID(payment *entity.Payment) (*model.Payment, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
    UPDATE payments
    SET 
        value_coin           = COALESCE($1, value_coin),
        value_forwarded_coin = COALESCE($2, value_forwarded_coin),
        currency             = COALESCE($3, currency),
        status               = COALESCE($4, status),
        address_in           = COALESCE($5, address_in),
        address_out          = COALESCE($6, address_out),
        txid_in              = COALESCE($7, txid_in),
        txid_out             = COALESCE($8, txid_out),
        confirmed_at         = COALESCE($9, confirmed_at)
    WHERE id = $10
    RETURNING
        id,
        user_id,
        value_coin,
        value_forwarded_coin,
        currency,
        status,
        address_in,
        address_out,
        txid_in,
        txid_out,
        created_at,
        confirmed_at
    `

	var response model.Payment

	err := r.db.QueryRow(
		query,
		payment.ValueCoin,
		payment.ValueForwardedCoin,
		payment.Currency,
		payment.Status,
		payment.AddressIn,
		payment.AddressOut,
		payment.TxidIn,
		payment.TxidOut,
		payment.ConfirmedAt,
		payment.ID,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.ValueCoin,
		&response.ValueForwardedCoin,
		&response.Currency,
		&response.Status,
		&response.AddressIn,
		&response.AddressOut,
		&response.TxidIn,
		&response.TxidOut,
		&response.ConfirmedAt,
		&response.CreatedAt,
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
		currency,
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
			&payment.Currency,
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
