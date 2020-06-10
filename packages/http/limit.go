package http

type Limit struct {
	Symbol string  `json:"symbol"`
	Limit  float64 `json:"limit"`
}
