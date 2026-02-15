package repository

import (
	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (r paymentRepositoryImpl) Create(payment *entity.Payment) (*model.Payment, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO payments (
			user_id,
			currency
		)
		VALUES ($1, $2)
		RETURNING
			id,
			currency
	`

	var response model.Payment

	err := r.db.QueryRow(
		query,
		payment.UserID,
		payment.Currency,
	).Scan(
		&response.ID,
		&response.Currency,
	)

	return &response, err
}

func (r *paymentRepositoryImpl) Payment(payment *entity.Payment) (*model.Payment, error) {
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

	return &response, err
}

func (r *paymentRepositoryImpl) Update(payment *entity.Payment) (*model.Payment, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		UPDATE payments
		SET 
			value_coin           = COALESCE(NULLIF($1::numeric, 0), value_coin),
			value_forwarded_coin = COALESCE(NULLIF($2::numeric, 0), value_forwarded_coin),
			currency             = COALESCE(NULLIF($3::text, ''), currency),
			status               = COALESCE(NULLIF($4::text, ''), status),
			address_in           = COALESCE(NULLIF($5::text, ''), address_in),
			address_out          = COALESCE(NULLIF($6::text, ''), address_out),
			txid_in              = COALESCE(NULLIF($7::text, ''), txid_in),
			txid_out             = COALESCE(NULLIF($8::text, ''), txid_out),
			confirmed_at         = COALESCE($9, confirmed_at)
		WHERE id = $10
		RETURNING
			id, user_id, value_coin, value_forwarded_coin, currency, 
			status, address_in, address_out, txid_in, txid_out, 
			created_at, confirmed_at
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
		&response.CreatedAt,
		&response.ConfirmedAt,
	)

	return &response, err
}

func (r *paymentRepositoryImpl) UserPayment(payment *entity.Payment) (*[]model.Payment, error) {
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
		var row model.Payment

		if err := rows.Scan(
			&row.ID,
			&row.UserID,
			&row.Currency,
			&row.Status,
			&row.CreatedAt,
			&row.ConfirmedAt,
		); err != nil {
			return nil, err
		}

		response = append(response, row)
	}

	return &response, nil
}

func (r *paymentRepositoryImpl) Delete(payment *entity.Payment) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	err := r.db.QueryRow(`DELETE FROM payments WHERE id = $1;`, payment.ID).Scan()

	if err != nil {
		return false, err
	}

	return true, nil
}
