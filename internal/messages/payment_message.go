package messages

const MessagePayment = `<b>🌊 Mint'AS - Rechargement</b>

👛 <b>Veuillez sélectionnez votre méthode de paiement</b>

<b><i>💡 En cas de problème ou question veuillez contacter @weee</i></b>`

// @params: %s = currency, %s = transaction id, %s = address_in, %s = minimum_transaction_coin, %s = priority
const MessagePaymentCurrency = `<b>🌊 Mint'AS - Rechargement (%s)</b>

ℹ️ <b>Votre paiement a été crée avec succès !</b>

 <b><code>•</code> Transaction ID: <code>%s</code></b>
 <b><code>•</code> Statut: <code>%s</code></b>
 <b><code>•</code> Adresse: <code>%s</code></b>
 <b><code>•</code> Dêpot minimum: <code>%s</code></b>
 <b><code>•</code> Priorité: <code>%s</code></b>

❗️ <b><i>(La validation du paiement peut prendre entre 5 à 30 minutes)</i></b>

<b><i>💡 En cas de problème ou question veuillez contacter @weee</i></b>`

const MessagePaymentConfirmed = `
<b>✅ Votre paiement de %v %s a été validé avec succès!</b>

 <b><code>•</code> Transaction ID: <code>%s</code></b>
 <b><code>•</code> Montant reçu: <code>%v€</code></b>
 <b><code>•</code> Créer le: <code>%v</code></b>
 <b><code>•</code> Confirmer le: <code>%v</code></b>

 <b><i>💡 En cas de problème ou question veuillez contacter @weee</i></b>`
