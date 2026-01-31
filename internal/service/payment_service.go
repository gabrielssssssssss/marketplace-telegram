package service

import (
	"context"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
	cryptapi "github.com/gabrielssssssssss/marketplace-telegram/libs/crypt-api"
	"github.com/go-telegram/bot/models"
)

type PaymentService interface {
	PaymentCurrencyCallback(ctx context.Context, callback *models.CallbackQuery) (*model.Payment, *cryptapi.PaymentResponse, error)
}

type paymentServiceImpl struct {
	repository repository.PaymentRepository
}

func NewPaymentService(paymentRepository repository.PaymentRepository) PaymentService {
	return &paymentServiceImpl{
		repository: paymentRepository,
	}
}
