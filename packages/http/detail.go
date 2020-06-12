package http

type Detail struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type DetailResponse struct {
	Symbol  string   `json:"symbol"`
	Details []Detail `json:"contracts"`
}
