package repository

import (
	"context"

	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (r accountRepositoryImpl) Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{ChatID: update.Message.Chat.ID, Text: "yo"})
}

func (r accountRepositoryImpl) CreateUser(users *entity.Users) (*model.Users, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO users (
			user_id,
			display_name,
			username,
			balance,
			recovery_key,
			updated_at,
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING
			"recovery_key"
	`

	var response model.Users

	err := r.db.QueryRow(
		query,
		users.UserId,
		users.DisplayName,
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

func (r accountRepositoryImpl) GetUserByID() {}

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
