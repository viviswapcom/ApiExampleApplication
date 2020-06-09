package order

import "github.com/tangleMesh/OmokuApiExampleApplication/packages/http"

type Order struct {
	CurrencyPair        http.CurrencyPair
	SourcePaymentMethod http.Method
	TargetPaymentMethod http.Method
	UserMail            string
	LoginCode           string
	SessionToken        string
	SessionSecret       string
	TwoFactor           bool
}
