package repository

import (
	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (r accountRepositoryImpl) CreateUser(users *entity.Users) (*model.Users, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO users (
			user_id,
			firstname,
			lastname,
			username,
			balance,
			recovery_key,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			"recovery_key"
	`

	var response model.Users

	err := r.db.QueryRow(
		query,
		users.UserId,
		users.Firstname,
		users.Lastname,
		users.Username,
		users.Balance,
		users.RecoveryKey,
		users.UpdatedAt,
	).Scan(&response.RecoveryKey)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r accountRepositoryImpl) GetUserByID(users *entity.Users) (*model.Users, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		user_id,
		firstname,
		lastname,
		username,
		balance,
		recovery_key,
		created_at,
		updated_at
	FROM users
	WHERE user_id = $1
	`

	var response model.Users

	err := r.db.QueryRow(
		query,
		users.UserId,
	).Scan(
		&response.UserId,
		&response.Firstname,
		&response.Lastname,
		&response.Username,
		&response.Balance,
		&response.RecoveryKey,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r accountRepositoryImpl) UpdateUser() {}

func (r accountRepositoryImpl) DeleteUser(users *entity.Users) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	_, err := r.db.Query(`DELETE FROM users WHERE user_id = $1;`, users.UserId)
	if err != nil {
		return false, err
	}

	return true, nil
}
