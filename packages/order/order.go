package order

import "github.com/tangleMesh/OmokuApiExampleApplication/packages/http"

type Order struct {
	ID                  string
	CurrencyPair        http.CurrencyPair
	SourcePaymentMethod http.Method
	TargetPaymentMethod http.Method
	UserMail            string
	LoginCode           string
	SessionToken        string
	SessionSecret       string
	TwoFactor           bool
	Amount              float64
	Connection          http.Connection
	SourcePaymentDetail http.Detail
	TargetPaymentDetail http.Detail
}
