package http

type Currency struct {
	Short          string `json:"short"`
	Name           string `json:"name"`
	Character      string `json:"character"`
	IsoCode        string `json:"isoCode"`
	IsDigitalAsset bool   `json:"isDigitalAsset"`
}

type CurrencyPair struct {
	Symbol         string   `json:"symbol"`
	SourceCurrency Currency `json:"sourceCurrency"`
	TargetCurrency Currency `json:"targetCurrency"`
}
