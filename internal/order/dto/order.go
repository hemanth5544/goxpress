package dto

type OrderMessage struct {
	Message string "json:message"
}

type CheckoutRequest struct {
	PaymentMethod string `json:"payment_method" binding:"required"`
}
