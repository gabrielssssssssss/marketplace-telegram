package service

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	cryptapi "github.com/gabrielssssssssss/marketplace-telegram/libs/crypt-api"
	"github.com/go-telegram/bot/models"
)

func (s *paymentServiceImpl) CreatePayment(ctx context.Context, callback *models.CallbackQuery) (*model.Payment, *cryptapi.PaymentResponse, error) {
	parts := strings.Split(callback.Data, "_")
	currency := parts[2]

	newPayment := &entity.Payment{
		UserID:   &callback.Message.Message.Chat.ID,
		Currency: &currency,
	}

	payment, err := s.repository.CreatePayment(newPayment)
	if err != nil {
		return nil, nil, err
	}

	hook := fmt.Sprintf("%s?payment_id=%s", os.Getenv("CALLBACK_URL"), payment.ID)
	client := cryptapi.NewCryptAPI("https://api.cryptapi.io/", hook)

	req := cryptapi.PaymentRequest{
		Address:  os.Getenv(strings.ToUpper(currency)),
		Currency: currency,
	}

	paymentInfo, err := client.CreatePayment(ctx, req)
	if err != nil {
		return payment, nil, err
	}

	return payment, paymentInfo, nil
}

func (s *paymentServiceImpl) FindPayment(payment *entity.PaymentCallback) (*model.Payment, error) {
	resp, err := s.repository.GetPaymentByID(payment.PaymentID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *paymentServiceImpl) ConfirmPayment(payment *entity.PaymentCallback) (*model.Payment, error) {
	valueCoin, err := strconv.ParseFloat(payment.ValueCoin, 64)
	if err != nil {
		return nil, err
	}

	valueForwardedCoin, err := strconv.ParseFloat(payment.ValueForwardedCoin, 64)
	if err != nil {
		return nil, err
	}

	updatePayment := entity.Payment{
		ID:                 &payment.PaymentID,
		AddressIn:          &payment.AddressIn,
		AddressOut:         &payment.AdddressOut,
		ValueCoin:          &valueCoin,
		ValueForwardedCoin: &valueForwardedCoin,
		TxidIn:             &payment.TxidIn,
		TxidOut:            &payment.TxidOut,
		ConfirmedAt:        &time.Time{},
		Status:             &payment.Status,
	}

	resp, err := s.repository.UpdatePaymentByID(&updatePayment)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
