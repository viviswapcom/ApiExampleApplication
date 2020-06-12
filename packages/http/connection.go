package http

type ConnectionDetail struct {
	PaymentMethod     Method      `json:"paymentMethod"`
	PaymentDetail     Detail      `json:"paymentDetail"`
	TransactionDetail interface{} `json:"transactionDetail"`
}

type Connection struct {
	ID     string           `json:"identifier"`
	Source ConnectionDetail `json:"source"`
	Target ConnectionDetail `json:"target"`
}

type ConnectionResponse struct {
	Symbol       string       `json:"symbol"`
	Connnections []Connection `json:"contracts"`
}
