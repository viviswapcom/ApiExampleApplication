package config

type Configuration struct {
	APIKey          string `json:"apiKey"`
	ApplicationName string `json:"applicationName"`
	Port            int    `json:"port"`
}
