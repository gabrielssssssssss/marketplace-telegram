package service

import (
	"context"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type PaymentService interface {
	PaymentCallback(ctx context.Context, b *bot.Bot, update *models.Update)
}

type paymentServiceImpl struct {
	repository repository.PaymentRepository
}

func NewPaymentService(paymentRepository repository.PaymentRepository) PaymentService {
	return &paymentServiceImpl{
		repository: paymentRepository,
	}
}
