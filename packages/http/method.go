package http

type Method struct {
	Key       string  `json:"key"`
	Name      string  `json:"name"`
	MinAmount float64 `json:"minAmount"`
	MaxAmount string  `json:"maxAmount"`
}

type MethodResponse struct {
	Symbol               string   `json:"symbol"`
	SourcePaymentMethods []Method `json:"sourcePaymentMethods"`
	TargetPaymentMethods []Method `json:"targetPaymentMethods"`
}
