package webhooks

import (
	"net/http"
)

func (webhook *PaymentWebhook) Webhooks(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/callback", webhook.WebhookPayment)
}
