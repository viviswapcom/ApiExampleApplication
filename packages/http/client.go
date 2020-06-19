package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tangleMesh/OmokuApiExampleApplication/packages/config"
)

func request(method string, url string, formData string, sessionToken string, sessionSecret string) ([]byte, Error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(formData)))
	config := config.GetConfig()
	req.Header.Set("X-API-KEY", config.APIKey)
	req.Header.Set("X-SESSION-TOKEN", sessionToken)
	req.Header.Set("X-SESSION-SECRET", sessionSecret)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	// Validate response
	if err != nil {
		log.Println(err)
		return nil, Error{
			Message: "Error doing the request with url: " + url,
		}
	}

	// Validate returned body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, Error{
			Message: "Error parsing the response body with url: " + url,
		}
	}

	// LOG response
	// fmt.Println("RESPONSE", url, resp, err, string(body))

	// Check if response is an error
	var errorVar Error = Error{}
	if string(body)[:8] == "{\"code\":" {
		json.Unmarshal([]byte(body), &errorVar)
		return nil, errorVar
	}

	// Check if response is a valid JSON
	return []byte(body), Error{}
	// json.Unmarshal([]byte(body), &jsonBody)
}

func GetCurrencyPairs() []CurrencyPair {
	var jsonBody []CurrencyPair
	resp, error := request("GET", "https://api-gateway-dev.omoku.io/currencies", "", "", "")

	if error != (Error{}) || resp == nil {
		log.Println(error)
		return nil
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody
}

func GetPaymentMethods(symbol string) MethodResponse {
	var jsonBody MethodResponse
	resp, error := request("GET", "https://api-gateway-dev.omoku.io/payment-methods/"+symbol, "", "", "")

	if error != (Error{}) || resp == nil {
		log.Println(error)
		return MethodResponse{}
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody
}

func GetLogin(mail string) (Login, Error) {
	var jsonBody Login
	resp, err := request("POST", "https://api-gateway-dev.omoku.io/login", "{\"mail\":\""+mail+"\"}", "", "")

	if err != (Error{}) || resp == nil {
		log.Println(err)
		return Login{}, err
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody, Error{}
}

func DoLogin(verificationCode string, sessionToken string) (LoginConfirmation, Error) {
	var jsonBody LoginConfirmation
	resp, err := request("POST", "https://api-gateway-dev.omoku.io/confirm-login", "{\"code\":\""+verificationCode+"\", \"sessionToken\":\""+sessionToken+"\"}", "", "")

	if err != (Error{}) || resp == nil {
		log.Println(err)
		return LoginConfirmation{}, err
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody, Error{}
}

func GetCourse(symbol string) (CourseResponse, Error) {
	var jsonBody CourseResponse
	resp, err := request("GET", "https://api-gateway-dev.omoku.io/courses/"+symbol, "", "", "")

	if err != (Error{}) || resp == nil {
		log.Println(err)
		return CourseResponse{}, err
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody, Error{}
}

func GetLimit(symbol string, sessionToken string, sessionSecret string) (Limit, Error) {
	var jsonBody Limit
	resp, err := request("GET", "https://api-gateway-dev.omoku.io/limits/"+symbol, "", sessionToken, sessionSecret)

	if err != (Error{}) || resp == nil {
		log.Println(err)
		return Limit{}, err
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody, Error{}
}

func GetDetails(currencyShort string, sessionToken string, sessionSecret string) (DetailResponse, Error) {
	var jsonBody DetailResponse
	resp, err := request("GET", "https://api-gateway-dev.omoku.io/payment-details/"+currencyShort, "", sessionToken, sessionSecret)

	if err != (Error{}) || resp == nil {
		log.Println(err)
		return DetailResponse{}, err
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody, Error{}
}

func GetConnections(symbol string, sessionToken string, sessionSecret string) (ConnectionResponse, Error) {
	var jsonBody ConnectionResponse
	resp, err := request("GET", "https://api-gateway-dev.omoku.io/connections/"+symbol, "", sessionToken, sessionSecret)

	if err != (Error{}) || resp == nil {
		log.Println(err)
		return ConnectionResponse{}, err
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody, Error{}
}

func CreateDetail(currencyShort string, address string, sessionToken string, sessionSecret string) (CreateDetailResponse, Error) {
	var jsonBody CreateDetailResponse
	resp, err := request("POST", "https://api-gateway-dev.omoku.io/payment-details/"+currencyShort, "{\"address\": \""+address+"\"}", sessionToken, sessionSecret)

	if err != (Error{}) || resp == nil {
		log.Println(err)
		return CreateDetailResponse{}, err
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody, Error{}
}

func InitOrder(symbol string, amount float64, sourcePaymentMethodKey string, sourcePaymentDetailId string, targetPaymentMethodKey string, targetPaymentDetailId string, sessionToken string, sessionSecret string) (OrderResponse, Error) {
	var jsonBody OrderResponse
	resp, err := request("POST", "https://api-gateway-dev.omoku.io/orders", "{\"symbol\": \""+symbol+"\", \"amount\": "+fmt.Sprintf("%g", amount)+", \"source\": { \"paymentMethod\": { \"key\": \""+sourcePaymentMethodKey+"\" }, \"paymentDetail\": { \"id\": \""+sourcePaymentDetailId+"\" }}, \"target\": { \"paymentMethod\": { \"key\": \""+targetPaymentMethodKey+"\" }, \"paymentDetail\": { \"id\": \""+targetPaymentDetailId+"\" }} }", sessionToken, sessionSecret)

	if err != (Error{}) || resp == nil {
		log.Println(err)
		return OrderResponse{}, err
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody, Error{}
}

func ConnfirmOrder(id string, confirmationCode string, sessionToken string, sessionSecret string) (ConfirmOrderResponse, Error) {
	var jsonBody ConfirmOrderResponse
	resp, err := request("POST", "https://api-gateway-dev.omoku.io/confirm-order", "{\"id\": \""+id+"\", \"code\": \""+confirmationCode+"\"}", sessionToken, sessionSecret)

	if err != (Error{}) || resp == nil {
		log.Println(err)
		return ConfirmOrderResponse{}, err
	}

	json.Unmarshal(resp, &jsonBody)
	return jsonBody, Error{}
}
