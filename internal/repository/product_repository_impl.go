package repository

import (
	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (r productRepositoryImpl) InsertProduct(product *entity.Product) (*model.Product, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO products (
			details,
			price
		)
		VALUES ($1, $2)
		RETURNING
			id,
			details,
			price,
			created_at,
			updated_at
	`

	var response model.Product

	err := r.db.QueryRow(
		query,
		product.Details,
		product.Price,
	).Scan(
		&response.ID,
		&response.Details,
		&response.Price,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	return &response, err
}

func (r productRepositoryImpl) SelectProductByID(product *entity.Product) (*model.Product, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		id,
		details,
		price,
		created_at,
		updated_at
	FROM products
	WHERE id = $1
	`

	var response model.Product

	err := r.db.QueryRow(
		query,
		product.ID,
	).Scan(
		&response.ID,
		&response.Details,
		&response.Price,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	return &response, err
}

func (r productRepositoryImpl) SelectAllProducts() (*[]model.Product, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		id,
		details,
		price,
		created_at,
		updated_at
	FROM products
	`

	var response []model.Product

	rows, err := r.db.Query(query)
	if err != nil {
		return &response, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		err = rows.Scan(
			&product.ID,
			&product.Details,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		response = append(response, product)
	}

	return &response, err
}

func (r *productRepositoryImpl) UpdateProductByID(product *entity.Product) (*model.Product, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
    UPDATE products
    SET 
		details      = COALESCE(NULLIF($1, '')::jsonb, details),
		price        = COALESCE($2, price),
		updated_at   = COALESCE($3, updated_at)
    WHERE id = $4
    RETURNING
		id,
		details,
		created_at,
		updated_at
    `

	var response model.Product

	err := r.db.QueryRow(
		query,
		product.Details,
		product.Price,
		product.UpdatedAt,
		product.ID,
	).Scan(
		&response.ID,
		&response.Details,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	return &response, err
}

func (r productRepositoryImpl) DeleteProductByID(product *entity.Product) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	_, err := r.db.Query(`DELETE FROM products WHERE id = $1;`, product.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
