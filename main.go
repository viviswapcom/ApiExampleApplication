package main

// import (
// 	"fyne.io/fyne"
// 	"fyne.io/fyne/app"
// 	"fyne.io/fyne/widget"

// 	"github.com/tangleMesh/OmokuApiExampleApplication/packages/config"
// )

// func main() {
// 	configuration := config.GetConfig()
// 	// currencyPairs := http.GetCurrencyPairs()

// 	a := app.New()

// 	w := a.NewWindow(configuration.ApplicationName)
// 	w.Resize(fyne.Size{600, 800})

// 	selectEntry := widget.NewSelectEntry([]string{"1", "2", "3"})
// 	w.SetContent(selectEntry)

// 	w.ShowAndRun()

// }

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/tangleMesh/OmokuApiExampleApplication/packages/config"

	"github.com/zserge/webview"
)

func main() {
	// Read config
	configuration := config.GetConfig()

	go startWebserver(configuration.Port)
	openWebview(true, configuration.ApplicationName, "http://localhost:"+strconv.Itoa(configuration.Port))
}

func startWebserver(port int) {
	fmt.Println("Server is listening on http://localhost:" + strconv.Itoa(port) + " or http://127.0.0.1:" + strconv.Itoa(port) + "")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := string(r.URL.Path)
		if path == "/" || path == "" {
			path = "index"
		}
		tmplFile, error := template.ParseFiles("./static/" + path + ".html")
		if error != nil {
			tmplFile, error = template.ParseFiles("./static/404.html")
		}

		tmpl := template.Must(tmplFile, error)
		tmpl.Execute(w, "") // , data
	})
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func openWebview(debug bool, appTitle string, webUrl string) {
	time.Sleep(2 * time.Second)
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle(appTitle)
	w.Bind("quit", func() {
		w.Terminate()
		os.Exit(3)
	})
	w.Navigate(webUrl)
	w.SetSize(600, 800, webview.HintNone)
	w.Run()
}
