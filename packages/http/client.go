package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tangleMesh/OmokuApiExampleApplication/packages/config"
)

func request(method string, url string) ([]byte, string) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequest(method, url, nil)
	config := config.GetConfig()
	req.Header.Set("X-API-KEY", config.APIKey)
	resp, err := client.Do(req)

	// Validate response
	if err != nil {
		log.Fatalln(err)
		return nil, "Error doing the request with url: " + url
	}

	// Validate returned body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, "Error parsing the response body with url: " + url
	}

	// Check if response is a valid JSON
	return []byte(body), ""
	// json.Unmarshal([]byte(body), &jsonBody)
}

func GetCurrencyPairs() []CurrencyPair {
	var jsonBody []CurrencyPair
	resp, error := request("GET", "https://api-gateway-dev.omoku.io/currencies")

	if error != "" || resp == nil {
		log.Fatalln(error)
		return nil
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody
}
