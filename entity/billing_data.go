package entity

type BillingData struct {
	CreateCustomer bool `json:"createCustomer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraudControl"`
	CheckoutPage   bool `json:"checkoutPage"`
}
