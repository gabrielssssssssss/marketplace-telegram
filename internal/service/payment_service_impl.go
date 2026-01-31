package service

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	cryptapi "github.com/gabrielssssssssss/marketplace-telegram/libs/crypt-api"
	"github.com/go-telegram/bot/models"
)

func (s *paymentServiceImpl) PaymentCurrencyCallback(ctx context.Context, callback *models.CallbackQuery) (*model.Payment, *cryptapi.PaymentResponse, error) {
	currency := strings.Split(callback.Data, "_")[2]

	payment := entity.Payment{
		UserID:   callback.Message.Message.Chat.ID,
		Currency: currency,
	}

	createdPayment, err := s.repository.CreatePayment(&payment)
	if err != nil {
		return nil, nil, err
	}

	callbackURL := fmt.Sprintf("%s?payment_id=%s", os.Getenv("CALLBACK_URL"), createdPayment.ID)
	cryptAPIClient := cryptapi.NewCryptAPI("https://api.cryptapi.io/", callbackURL)

	providerResponse, err := cryptAPIClient.CreatePayment(ctx, cryptapi.PaymentRequest{
		Address:  os.Getenv(strings.ToUpper(currency)),
		Currency: currency,
	})
	if err != nil {
		return createdPayment, nil, err
	}

	return createdPayment, providerResponse, nil
}
