package http

type Detail struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

type DetailResponse struct {
	CurrencyShort string   `json:"currencyShort"`
	Details       []Detail `json:"paymentDetails"`
}

type CreateDetailResponse struct {
	CurrencyShort string `json:"currencyShort"`
	Detail        Detail `json:"paymentDetail"`
}
