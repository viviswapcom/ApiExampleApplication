package http

type OrderDetail struct {
	PaymentMethod     Method      `json:"paymentMethod"`
	PaymentDetail     Detail      `json:"paymentDetail"`
	TransactionDetail interface{} `json:"transactionDetail"`
}

type OrderResponse struct {
	ID     string      `json:"id"`
	Symbol string      `json:"symbol"`
	Amount float32     `json:"amount"`
	Source OrderDetail `json:"source"`
	Target OrderDetail `json:"target"`
}

type ConfirmOrderResponse struct {
	ID      string `json:"id"`
	Success bool   `json:"success"`
}
