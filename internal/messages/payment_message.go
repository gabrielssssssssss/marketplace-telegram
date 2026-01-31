package messages

const MessagePayment = `<b>🌊 Mint'AS - Rechargement</b>

👛 <b>Veuillez sélectionnez votre méthode de paiement</b>

<b><i>💡 En cas de problème ou question veuillez contacter @weee</i></b>`

type PaymentResponse struct {
	Status                 string `json:"status"`
	AddressIn              string `json:"address_in"`
	AddressOut             string `json:"address_out"`
	CallbackUrl            string `json:"callback_url"`
	Priority               string `json:"priority"`
	MinimumTransactionCoin string `json:"minimum_transaction_coin"`
}

const MessagePaymentCurrency = `<b>🌊 Mint'AS - Rechargement (%s)</b>

✅ <b>Votre paiement a été crée avec succès !</b>

 <b><code>•</code> Statut: <code>%s</code></b>
 <b><code>•</code> Adresse: <code>%s</code></b>
 <b><code>•</code> Dêpot minimum: <code>%s</code></b>
 <b><code>•</code> Priorité: <code>%s</code></b>

<b><i>💡 En cas de problème ou question veuillez contacter @weee</i></b>`
