package repository

import (
	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (r userRepositoryImpl) CreateUser(user *entity.User) (*model.User, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO users (
			user_id,
			firstname,
			lastname,
			username,
			role,
			balance,
			recovery_key,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING
			"recovery_key"
	`

	var response model.User

	err := r.db.QueryRow(
		query,
		user.UserId,
		user.Firstname,
		user.Lastname,
		user.Username,
		user.Role,
		user.Balance,
		user.RecoveryKey,
		user.UpdatedAt,
	).Scan(&response.RecoveryKey)

	return &response, err
}

func (r userRepositoryImpl) GetUserByID(user *entity.User) (*model.User, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		user_id,
		firstname,
		lastname,
		username,
		role,
		balance,
		recovery_key,
		created_at,
		updated_at
	FROM users
	WHERE user_id = $1
	`

	var response model.User

	err := r.db.QueryRow(
		query,
		user.UserId,
	).Scan(
		&response.UserId,
		&response.Firstname,
		&response.Lastname,
		&response.Username,
		&response.Role,
		&response.Balance,
		&response.RecoveryKey,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	return &response, err
}

func (r userRepositoryImpl) GetUserByRecoveryKey(user *entity.User) (*model.User, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	SELECT
		user_id,
		firstname,
		lastname,
		username,
		role,
		balance,
		recovery_key,
		created_at,
		updated_at
	FROM users
	WHERE recovery_key = $1
	`

	var response model.User

	err := r.db.QueryRow(
		query,
		user.RecoveryKey,
	).Scan(
		&response.UserId,
		&response.Firstname,
		&response.Lastname,
		&response.Username,
		&response.Role,
		&response.Balance,
		&response.RecoveryKey,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	return &response, err
}

func (r *userRepositoryImpl) UpdateUserByID(user *entity.User) (*model.User, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
    UPDATE users
    SET 
		username     = COALESCE($1, username),
		firstname    = COALESCE($2, firstname),
		lastname     = COALESCE($3, lastname),
		balance      = COALESCE($4, balance),
		role         = COALESCE($5, role),
		recovery_key = COALESCE($6, recovery_key),
		updated_at   = COALESCE($7, updated_at)
    WHERE user_id = $8
    RETURNING
		user_id,
		username,
		firstname,
		lastname,
		role,
		balance,
		recovery_key,
		created_at,
		updated_at
    `

	var response model.User

	err := r.db.QueryRow(
		query,
		user.Username,
		user.Firstname,
		user.Lastname,
		user.Balance,
		user.Role,
		user.RecoveryKey,
		user.UpdatedAt,
		user.UserId,
	).Scan(
		&response.UserId,
		&response.Username,
		&response.Firstname,
		&response.Lastname,
		&response.Role,
		&response.Balance,
		&response.RecoveryKey,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	return &response, err
}

func (r userRepositoryImpl) DeleteUser(user *entity.User) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	_, err := r.db.Query(`DELETE FROM users WHERE user_id = $1;`, user.UserId)
	if err != nil {
		return false, err
	}

	return true, nil
}
