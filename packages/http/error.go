package http

type ErrorDetails struct {
	ErrorMessage string `json:"errorMessage"`
}

type Error struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Details ErrorDetails `json:"details"`
}
