package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/tangleMesh/OmokuApiExampleApplication/packages/config"
	omokuClient "github.com/tangleMesh/OmokuApiExampleApplication/packages/http"
	"github.com/tangleMesh/OmokuApiExampleApplication/packages/order"

	"github.com/zserge/webview"
)

func main() {
	// Read config
	configuration := config.GetConfig()

	// Trading configuration
	order := order.Order{}

	go startWebserver(configuration.Port)
	openWebview(true, configuration.ApplicationName, "http://localhost:"+strconv.Itoa(configuration.Port), &order)
}

func startWebserver(port int) {
	fmt.Println("Server is listening on http://localhost:" + strconv.Itoa(port) + " or http://127.0.0.1:" + strconv.Itoa(port) + "")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplFile, error := template.ParseFiles("./static/index.html")
		if error != nil {
			tmplFile, error = template.ParseFiles("./static/404.html")
		}
		data := omokuClient.GetCurrencyPairs()
		tmpl := template.Must(tmplFile, error)
		tmpl.Execute(w, data) // , data
	})
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func openWebview(debug bool, appTitle string, webUrl string, order *order.Order) {
	var methods omokuClient.MethodResponse
	time.Sleep(2 * time.Second)
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle(appTitle)
	w.Bind("setSymbol", func(symbol string) omokuClient.MethodResponse {
		fmt.Println("Selected Symbol:", symbol)

		// Save currency-pair
		currencyPairs := omokuClient.GetCurrencyPairs()
		for _, v := range currencyPairs {
			if v.Symbol == symbol {
				order.CurrencyPair = v
			}
		}

		// Load payment-methods
		methods = omokuClient.GetPaymentMethods(order.CurrencyPair.Symbol)
		return methods
	})
	w.Bind("setPaymentMethod", func(paymentMethodKey string) omokuClient.Method {
		fmt.Println("Selected Payment-method:", paymentMethodKey)

		for _, v := range methods.SourcePaymentMethods {
			if v.Key == paymentMethodKey {
				order.SourcePaymentMethod = v
			}
		}
		order.TargetPaymentMethod = methods.TargetPaymentMethods[0]

		return order.SourcePaymentMethod
	})
	w.Bind("initializeSession", func(mail string) url.Values {
		fmt.Println("Input login mail:", mail)

		login, err := omokuClient.GetLogin(mail)
		success := "false"
		if login.Success {
			success = "true"
			order.SessionToken = login.SessionToken
		}
		return url.Values{
			"err":   {err.Message},
			"login": {success},
		}
	})
	w.Bind("confirmSession", func(verificationCode string) url.Values {
		fmt.Println("Input verification code:", verificationCode)

		loginConf, err := omokuClient.DoLogin(verificationCode, order.SessionToken)
		success := "false"
		if loginConf.Success {
			success = "true"
			order.SessionSecret = loginConf.SessionSecret
		}

		return url.Values{
			"err":       {err.Message},
			"loginConf": {success},
		}
	})
	w.Bind("getCourse", func() omokuClient.CourseResponse {
		course, _ := omokuClient.GetCourse(order.CurrencyPair.Symbol)
		return course
	})
	w.Bind("getLimit", func() omokuClient.Limit {
		limit, _ := omokuClient.GetLimit(order.CurrencyPair.Symbol, order.SessionToken, order.SessionSecret)
		return limit
	})
	w.Bind("setAmount", func(amount float64) {
		fmt.Println("Input amount:", amount)
		order.Amount = amount
	})
	w.Bind("getDetails", func() omokuClient.DetailResponse {
		details, _ := omokuClient.GetDetails(order.CurrencyPair.SourceCurrency.Short, order.SessionToken, order.SessionSecret)
		return details
	})
	w.Bind("getConnections", func() omokuClient.ConnectionResponse {
		connections, _ := omokuClient.GetConnections(order.CurrencyPair.Symbol, order.SessionToken, order.SessionSecret)
		return connections
	})
	w.Bind("setConnection", func(connectionIdentifier string) {
		fmt.Println("Input connection:", connectionIdentifier)

		// Save connection
		connections, _ := omokuClient.GetConnections(order.CurrencyPair.Symbol, order.SessionToken, order.SessionSecret)
		for _, v := range connections.Connnections {
			if v.ID == connectionIdentifier {
				order.Connection = v
			}
		}
	})
	w.Bind("setSourcePaymentDetail", func(paymentDetailIdentifier string) omokuClient.DetailResponse {
		fmt.Println("Input source-payment-detail:", paymentDetailIdentifier)

		// Save source payment-detail
		details, _ := omokuClient.GetDetails(order.CurrencyPair.SourceCurrency.Short, order.SessionToken, order.SessionSecret)
		for _, v := range details.Details {
			if v.ID == paymentDetailIdentifier {
				order.SourcePaymentDetail = v
			}
		}

		targetDetails, _ := omokuClient.GetDetails(order.CurrencyPair.TargetCurrency.Short, order.SessionToken, order.SessionSecret)
		return targetDetails
	})
	w.Bind("setTargetPaymentDetail", func(paymentDetailIdentifier string) {
		fmt.Println("Input target-payment-detail:", paymentDetailIdentifier)

		// Save source payment-detail
		details, _ := omokuClient.GetDetails(order.CurrencyPair.TargetCurrency.Short, order.SessionToken, order.SessionSecret)
		for _, v := range details.Details {
			if v.ID == paymentDetailIdentifier {
				order.TargetPaymentDetail = v
			}
		}
	})
	w.Bind("createDetail", func(isTarget bool, address string) url.Values {
		fmt.Println("Input new address:", address, isTarget)

		currencyShort := order.CurrencyPair.SourceCurrency.Short
		if isTarget {
			currencyShort = order.CurrencyPair.TargetCurrency.Short
		}

		newDetailResponse, err := omokuClient.CreateDetail(currencyShort, address, order.SessionToken, order.SessionSecret)
		success := "false"
		if newDetailResponse.Detail.ID != "" {
			success = "true"
		}

		return url.Values{
			"err":        {err.Message},
			"success":    {success},
			"identifier": {newDetailResponse.Detail.ID},
		}
	})
	w.Bind("getOrderDetails", func() interface{} {
		return order
	})
	w.Bind("initializeNewOrder", func() url.Values {
		result, err := omokuClient.InitOrder(order.CurrencyPair.Symbol, order.Amount, order.SourcePaymentMethod.Key, order.SourcePaymentDetail.ID, order.TargetPaymentMethod.Key, order.TargetPaymentDetail.ID, order.SessionToken, order.SessionSecret)
		success := "false"
		if result.ID != "" {
			success = "true"
			order.ID = result.ID
		}

		return url.Values{
			"err":     {err.Message},
			"success": {success},
			"orderId": {order.ID},
		}
	})
	w.Bind("confirmNewOrder", func(confirmationCode string) url.Values {
		fmt.Println("Input confirmation code:", confirmationCode)

		result, err := omokuClient.ConnfirmOrder(order.ID, confirmationCode, order.SessionToken, order.SessionSecret)
		success := "false"
		if result.Success {
			success = "true"
		}

		return url.Values{
			"err":     {err.Message},
			"success": {success},
		}
	})

	w.Navigate(webUrl)
	w.SetSize(600, 800, webview.HintNone)
	w.Run()
}
