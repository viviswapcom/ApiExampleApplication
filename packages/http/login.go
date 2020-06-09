package http

type Login struct {
	Success      bool   `json:"success"`
	TwoFactor    bool   `json:"twoFactor"`
	SessionToken string `json:"sessionToken"`
}

type LoginConfirmation struct {
	Success       bool   `json:"success"`
	SessionSecret string `json:"sessionSecret"`
}
